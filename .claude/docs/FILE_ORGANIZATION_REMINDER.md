# Claude File Organization Reminder

## Important: Always use .claude/ directory structure

As defined in CLAUDE.md, all Claude-generated files must be organized in the `.claude/` directory:

```
.claude/
├── docs/           # Documentation and architectural guides
├── scripts/        # Automation scripts and build tools
├── coverage/       # Test coverage reports (gitignored)
└── archive/        # Historical/deprecated files
```

### Common Mistakes to Avoid
- ❌ Don't place scripts in root directory
- ❌ Don't create documentation files in root
- ❌ Don't leave temporary files scattered

### Correct Placement
- ✅ Python scripts: `.claude/scripts/fix_*.py`
- ✅ Shell scripts: `.claude/scripts/*.sh`
- ✅ Documentation: `.claude/docs/*.md`
- ✅ Coverage reports: `.claude/coverage/`
- ✅ Old/temp files: `.claude/archive/`

### Examples
- `fix_encoder_close.sh` → `.claude/scripts/fix_encoder_close.sh`
- `MIGRATION_STATUS.md` → `.claude/docs/MIGRATION_STATUS.md`
- `apply_xml_template.py` → `.claude/scripts/apply_xml_template.py`