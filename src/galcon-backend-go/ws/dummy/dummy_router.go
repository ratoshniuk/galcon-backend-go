package dummy

import "app"

var Routes = []*app.WSEndpoint{
	{
		URL:     "/echo",
		Handler: echo,
	},
	{
		URL:     "/",
		Handler: home,
	},
}
