# Repository Instructions

## Release Check

After each new commit in this repository, explicitly evaluate whether the change should trigger a new release.

Use a release when the commit introduces user-facing changes such as:
- public header/API changes
- behavior fixes that downstream users depend on
- compatibility changes
- build, packaging, or installation changes that affect consumers
- any change that should be reflected in `CHANGELOG.md`

If the commit is internal-only and does not affect downstream users, a release is usually not required. Make the release decision explicitly instead of skipping the check.

## Post-Release Follow-up

Whenever a new `limonp` version is released, create or update an issue in `yanyiwu/cppjieba` asking it to update to the latest `limonp` version.

The issue should include:
- the new `limonp` version
- the release/tag link
- a short summary of the downstream-relevant changes
