# Claude Generated Files

This directory contains files generated during development with Claude Code. These files are organized for better project structure and maintainability.

## Directory Structure

### 📚 `docs/` - Documentation
Contains architectural documentation, design decisions, and implementation guides generated during development:

- **BASE_ABSTRACTIONS.md** - Architecture documentation for base abstractions pattern
- **ERROR_DESIGN_PROPOSAL.md** - Error handling design specifications
- **IMPLEMENTATION_GUIDE.md** - Step-by-step guide for adding new message types
- **MIGRATION_PATTERNS.md** - Patterns used during architecture migrations
- **MIGRATION_STATUS.md** - Status tracking for migration tasks
- **TYPE_SAFETY_ANALYSIS.md** - Analysis of type safety improvements
- **TYPE_SAFETY_RESULTS.md** - Results from type safety implementation
- **WRAPPER_MIGRATION_PLAN.md** - Plan for migrating wrapper package
- **XML_TO_GO_MAPPING.md** - Critical guide for XML field mapping
- **path-remapping-inconsistencies.md** - Documentation of path mapping issues
- **refactoring.md** - Refactoring notes and decisions
- **test-coverage-strategy.md** - Testing strategy documentation

### 🔧 `scripts/` - Automation Scripts
Contains Python and shell scripts used for code generation, refactoring, and automation:

- **apply_idiomatic_xml_api.py** - Script to apply XML-first API patterns
- **apply_xml_template.py** - Template application for XML methods
- **batch_apply_xml_api.sh** - Batch processing for XML API updates
- **complete_refactoring.sh** - Complete refactoring automation
- **complete_remaining_xml_api.sh** - Finalize XML API implementation
- **fix_*.py/sh** - Various fix and cleanup scripts for specific issues
- **refactor_to_idiomatic*.sh** - Refactoring scripts for idiomatic patterns
- **remove_message_with.py** - Script to remove deprecated MessageWith usage

### 📊 `coverage/` - Test Coverage Reports
Contains test coverage output files:

- **cover.out** - Go coverage output
- **coverage.out** - Additional coverage data
- **coverage.txt** - Coverage report in text format

### 📦 `archive/` - Historical Files
Contains deprecated or temporary files kept for reference:

- **CustomerCreditTransfer.test** - Test output file
- **gitleaks.tar.gz** - Security scanning archive

## Usage

### For Developers
- **docs/**: Reference these files when working on architecture or adding new features
- **scripts/**: Run these scripts when needed for automation tasks
- **coverage/**: Review coverage reports to understand test coverage
- **archive/**: Historical reference only

### For Maintenance
Most files in `scripts/` and `archive/` can be safely removed after the project is stable. The `docs/` directory should be preserved as it contains important architectural decisions and implementation guides.

## Best Practices

1. **Keep docs/ updated** - When architecture changes, update relevant documentation
2. **Clean scripts/ periodically** - Remove obsolete automation scripts
3. **Archive old coverage** - Keep recent coverage reports, archive old ones
4. **Git ignore coverage/** - Coverage files shouldn't be committed to repository

---

*This directory structure follows best practices for keeping project root clean while preserving important development artifacts.*