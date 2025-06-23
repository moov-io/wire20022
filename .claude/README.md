# .claude Directory

This directory contains all Claude-generated files for the wire20022 project.

## Directory Structure

```
.claude/
├── docs/           # Documentation and architectural guides
├── scripts/        # Automation scripts and build tools
├── coverage/       # Test coverage reports (gitignored)
└── archive/        # Historical/deprecated files
```

## Purpose

This structure keeps Claude-generated files organized and separate from the main project code, ensuring:
- Clean project root directory
- Easy identification of AI-generated content
- Proper gitignore rules for temporary files
- Clear separation between project code and development tools

## Guidelines

All Claude-generated content should be placed in the appropriate subdirectory:
- Documentation and guides → `docs/`
- Scripts and automation → `scripts/`
- Coverage reports → `coverage/` (automatically gitignored)
- Old or temporary files → `archive/`