#!/usr/bin/env python3
"""
Field Evolution Analyzer for wire20022 Message Types

Analyzes map.go files to identify version-specific field patterns
and outputs field grouping recommendations for type-safe migration.

Usage: python3 analyze_field_evolution.py [message_type]
"""

import os
import re
import json
from pathlib import Path
from typing import Dict, List, Set, Tuple
from collections import defaultdict

class FieldEvolutionAnalyzer:
    def __init__(self, repo_root: str = "."):
        self.repo_root = Path(repo_root)
        self.models_dir = self.repo_root / "pkg" / "models"
        
    def find_message_types(self) -> List[str]:
        """Find all message type directories in pkg/models/"""
        message_types = []
        for item in self.models_dir.iterdir():
            if item.is_dir() and (item / "map.go").exists():
                message_types.append(item.name)
        return sorted(message_types)
    
    def parse_map_file(self, message_type: str) -> Dict[str, Dict[str, str]]:
        """Parse map.go file to extract version-specific field mappings"""
        map_file = self.models_dir / message_type / "map.go"
        
        if not map_file.exists():
            return {}
            
        with open(map_file, 'r') as f:
            content = f.read()
        
        # Extract PathMapVX functions
        version_mappings = {}
        
        # Pattern to match function definitions
        func_pattern = r'func PathMapV(\d+)\(\) map\[string\]any \{([^}]+(?:\{[^}]*\}[^}]*)*)\}'
        
        matches = re.finditer(func_pattern, content, re.DOTALL)
        
        for match in matches:
            version = match.group(1)
            body = match.group(2)
            
            # Parse field mappings from function body
            field_mappings = self.parse_field_mappings(body)
            version_mappings[f"V{version}"] = field_mappings
            
        return version_mappings
    
    def parse_field_mappings(self, body: str) -> Dict[str, str]:
        """Parse field mappings from PathMapVX function body"""
        mappings = {}
        
        # Simple string mappings: "xml.path": "go.field"
        simple_pattern = r'"([^"]+)":\s*"([^"]+)"'
        
        for match in re.finditer(simple_pattern, body):
            xml_path = match.group(1)
            go_field = match.group(2)
            mappings[xml_path] = go_field
            
        return mappings
    
    def analyze_field_evolution(self, message_type: str) -> Dict:
        """Analyze how fields evolve across versions for a message type"""
        version_mappings = self.parse_map_file(message_type)
        
        if not version_mappings:
            return {"error": f"No mappings found for {message_type}"}
        
        # Track when each field was introduced
        field_introduction = {}
        all_fields = set()
        
        # Collect all fields across all versions
        for version, mappings in version_mappings.items():
            for go_field in mappings.values():
                all_fields.add(go_field)
        
        # Determine introduction version for each field
        for field in all_fields:
            for version in sorted(version_mappings.keys(), key=lambda x: int(x[1:])):
                if any(field == go_field for go_field in version_mappings[version].values()):
                    if field not in field_introduction:
                        field_introduction[field] = version
                    break
        
        # Group fields by introduction version
        version_groups = defaultdict(list)
        for field, version in field_introduction.items():
            version_groups[version].append(field)
        
        # Analyze field patterns
        patterns = self.identify_field_patterns(field_introduction, version_mappings)
        
        return {
            "message_type": message_type,
            "total_versions": len(version_mappings),
            "total_fields": len(all_fields),
            "field_introduction": field_introduction,
            "version_groups": dict(version_groups),
            "patterns": patterns,
            "recommendations": self.generate_recommendations(patterns, version_groups)
        }
    
    def identify_field_patterns(self, field_introduction: Dict[str, str], 
                               version_mappings: Dict[str, Dict[str, str]]) -> Dict:
        """Identify common patterns in field evolution"""
        patterns = {
            "business_query_fields": [],
            "reporting_fields": [],
            "address_fields": [],
            "transaction_fields": [],
            "date_fields": [],
            "agent_fields": [],
            "core_fields": []
        }
        
        for field, version in field_introduction.items():
            field_lower = field.lower()
            
            # Business query pattern
            if any(term in field_lower for term in ['business', 'query', 'bizqry', 'orgnlbizqry']):
                patterns["business_query_fields"].append({"field": field, "version": version})
            
            # Reporting pattern  
            elif any(term in field_lower for term in ['reporting', 'sequence', 'rptgseq']):
                patterns["reporting_fields"].append({"field": field, "version": version})
            
            # Address pattern
            elif any(term in field_lower for term in ['address', 'floor', 'room', 'building', 'bldg']):
                patterns["address_fields"].append({"field": field, "version": version})
            
            # Transaction pattern
            elif any(term in field_lower for term in ['uetr', 'transaction', 'unique', 'ref']):
                patterns["transaction_fields"].append({"field": field, "version": version})
            
            # Date pattern
            elif any(term in field_lower for term in ['date', 'datetime', 'dt']):
                patterns["date_fields"].append({"field": field, "version": version})
            
            # Agent pattern
            elif any(term in field_lower for term in ['agent', 'instg', 'instd', 'intrmyagt']):
                patterns["agent_fields"].append({"field": field, "version": version})
            
            # Core fields (introduced early)
            elif version in ['V1', 'V2', 'V3']:
                patterns["core_fields"].append({"field": field, "version": version})
        
        return patterns
    
    def generate_recommendations(self, patterns: Dict, version_groups: Dict) -> Dict:
        """Generate migration recommendations based on patterns"""
        recommendations = {
            "migration_pattern": "unknown",
            "field_groups": [],
            "complexity": "low",
            "priority": "low"
        }
        
        # Determine migration pattern
        if patterns["business_query_fields"] and patterns["reporting_fields"]:
            recommendations["migration_pattern"] = "reporting_evolution"
            recommendations["field_groups"] = [
                {
                    "name": "BusinessQueryFields",
                    "fields": [f["field"] for f in patterns["business_query_fields"]],
                    "introduced": "V3+"
                },
                {
                    "name": "ReportingFields", 
                    "fields": [f["field"] for f in patterns["reporting_fields"]],
                    "introduced": "V7+"
                }
            ]
        
        elif patterns["address_fields"] or patterns["transaction_fields"]:
            recommendations["migration_pattern"] = "payment_evolution"
            if patterns["address_fields"]:
                recommendations["field_groups"].append({
                    "name": "AddressEnhancementFields",
                    "fields": [f["field"] for f in patterns["address_fields"]],
                    "introduced": "V8+"
                })
            if patterns["transaction_fields"]:
                recommendations["field_groups"].append({
                    "name": "TransactionFields",
                    "fields": [f["field"] for f in patterns["transaction_fields"]],
                    "introduced": "V7+"
                })
        
        elif len(version_groups) <= 3:
            recommendations["migration_pattern"] = "simple_stable"
            recommendations["field_groups"] = [
                {
                    "name": "CoreFields",
                    "fields": list({f for fields in version_groups.values() for f in fields}),
                    "introduced": "V1+"
                }
            ]
        
        # Determine complexity
        total_versions = len(version_groups)
        total_field_groups = len([g for g in recommendations["field_groups"] if g["fields"]])
        
        if total_versions >= 10 and total_field_groups >= 2:
            recommendations["complexity"] = "high"
            recommendations["priority"] = "high"
        elif total_versions >= 6 or total_field_groups >= 2:
            recommendations["complexity"] = "medium"
            recommendations["priority"] = "medium"
        else:
            recommendations["complexity"] = "low"
            recommendations["priority"] = "low"
            
        return recommendations
    
    def analyze_all_message_types(self) -> Dict:
        """Analyze all message types and generate comprehensive report"""
        message_types = self.find_message_types()
        results = {}
        
        for message_type in message_types:
            print(f"Analyzing {message_type}...")
            results[message_type] = self.analyze_field_evolution(message_type)
        
        # Generate summary
        summary = self.generate_summary(results)
        
        return {
            "summary": summary,
            "message_types": results
        }
    
    def generate_summary(self, results: Dict) -> Dict:
        """Generate summary of all analysis results"""
        total_messages = len(results)
        complexity_counts = defaultdict(int)
        pattern_counts = defaultdict(int)
        priority_counts = defaultdict(int)
        
        for message_type, analysis in results.items():
            if "error" in analysis:
                continue
                
            recs = analysis.get("recommendations", {})
            complexity_counts[recs.get("complexity", "unknown")] += 1
            pattern_counts[recs.get("migration_pattern", "unknown")] += 1
            priority_counts[recs.get("priority", "unknown")] += 1
        
        return {
            "total_message_types": total_messages,
            "complexity_distribution": dict(complexity_counts),
            "pattern_distribution": dict(pattern_counts),
            "priority_distribution": dict(priority_counts),
            "migration_order": self.suggest_migration_order(results)
        }
    
    def suggest_migration_order(self, results: Dict) -> List[str]:
        """Suggest optimal migration order based on complexity and priority"""
        # Priority order: high complexity + high priority first
        message_priorities = []
        
        for message_type, analysis in results.items():
            if "error" in analysis:
                continue
                
            recs = analysis.get("recommendations", {})
            complexity = recs.get("complexity", "low")
            priority = recs.get("priority", "low")
            
            # Scoring system
            score = 0
            if priority == "high": score += 10
            elif priority == "medium": score += 5
            
            if complexity == "high": score += 3
            elif complexity == "medium": score += 1
            
            message_priorities.append((message_type, score))
        
        # Sort by score (descending) and return message types
        message_priorities.sort(key=lambda x: x[1], reverse=True)
        return [msg for msg, score in message_priorities]

def main():
    import argparse
    
    parser = argparse.ArgumentParser(description="Analyze field evolution for wire20022 message types")
    parser.add_argument("message_type", nargs="?", help="Specific message type to analyze")
    parser.add_argument("--all", action="store_true", help="Analyze all message types")
    parser.add_argument("--output", help="Output file for results (JSON)")
    
    args = parser.parse_args()
    
    analyzer = FieldEvolutionAnalyzer()
    
    if args.all or not args.message_type:
        results = analyzer.analyze_all_message_types()
        
        # Print summary
        print("\n" + "="*60)
        print("FIELD EVOLUTION ANALYSIS SUMMARY")
        print("="*60)
        
        summary = results["summary"]
        print(f"Total message types analyzed: {summary['total_message_types']}")
        print(f"Complexity distribution: {summary['complexity_distribution']}")
        print(f"Pattern distribution: {summary['pattern_distribution']}")
        print(f"Priority distribution: {summary['priority_distribution']}")
        
        print(f"\nSuggested migration order:")
        for i, msg_type in enumerate(summary['migration_order'][:10], 1):
            analysis = results["message_types"][msg_type]
            recs = analysis.get("recommendations", {})
            print(f"{i:2d}. {msg_type:25s} ({recs.get('complexity', 'unknown'):6s} complexity, {recs.get('migration_pattern', 'unknown')})")
        
    else:
        results = analyzer.analyze_field_evolution(args.message_type)
        
        # Print detailed analysis
        print(f"\nAnalysis for {args.message_type}:")
        print(f"Total versions: {results.get('total_versions', 0)}")
        print(f"Total fields: {results.get('total_fields', 0)}")
        
        recommendations = results.get("recommendations", {})
        print(f"Migration pattern: {recommendations.get('migration_pattern', 'unknown')}")
        print(f"Complexity: {recommendations.get('complexity', 'unknown')}")
        print(f"Priority: {recommendations.get('priority', 'unknown')}")
        
        if recommendations.get("field_groups"):
            print(f"\nRecommended field groups:")
            for group in recommendations["field_groups"]:
                print(f"  {group['name']} ({group['introduced']}): {len(group['fields'])} fields")
    
    # Save results if output file specified
    if args.output:
        with open(args.output, 'w') as f:
            json.dump(results, f, indent=2)
        print(f"\nResults saved to {args.output}")

if __name__ == "__main__":
    main()