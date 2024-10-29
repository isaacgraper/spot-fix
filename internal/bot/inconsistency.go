package bot

func (pr *Process) VerifiyInconsistence() bool {
	return len(pr.Results) > 0
}
