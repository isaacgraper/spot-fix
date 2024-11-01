package page

func (p *Page) Loading() {
	p.Rod.MustWaitLoad().MustWaitStable().MustWaitDOMStable()
}
