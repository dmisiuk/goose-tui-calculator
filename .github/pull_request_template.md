## Summary
(What does this change do?)

Closes #ISSUE_ID

## Before
(Embed GIF or write `N/A`)

## After
(Embed GIF)

> **🎬 Automated Demos**: The VHS Demo Recording workflow will automatically generate demos for this PR.
> Check the Actions tab or wait for the bot comment with download links. You can download and commit
> the generated GIFs from artifacts if they meet quality standards.

## Updated / Added Tapes
- (List each changed or new .tape)

## Tests
- [ ] Added/Updated unit tests
- [ ] Not applicable (explain why)

## Checklist
- [ ] Issue referenced (Closes #ID)
- [ ] Branch follows naming convention
- [ ] VHS tape(s) updated (or N/A with justification)
- [ ] GIF(s) regenerated in .tapes/assets/ (or downloaded from automated workflow artifacts)
- [ ] go vet + go fmt clean
- [ ] go test ./... passes
- [ ] README/docs updated if needed
- [ ] No stray debug prints
- [ ] go mod tidy produces no diff

### Demo Quality Checklist (if using automated demos)
- [ ] Downloaded generated demos from workflow artifacts
- [ ] Verified demo accuracy and visual quality
- [ ] Confirmed timing and interactions are smooth
- [ ] Committed approved demos to `.tapes/assets/`
