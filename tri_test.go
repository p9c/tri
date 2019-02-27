package tri

var exampleTri = Tri{
	"appname",
	Brief{"brief"},
	Version{0, 1, 1, "alpha"},
	DefaultCommand{"help"},
	Var{"datadir",
		Short{"d"},
		Brief{"brief"},
		Usage{"usage"},
		Help{"help"},
		Default{"~/.pod"},
		Slot{""},
	},
	Commands{
		{"ctl",
			Short{"c"},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Group{"groupname"},
			Examples{
				"example 1", "explaining text",
				"example 2", "explaining text",
			},
			Var{"datadir",
				Short{"d"},
				Brief{"brief"},
				Usage{"usage"},
				Help{"help"},
				Default{"~/.pod"},
				Slot{""},
			},
			Trigger{"init",
				Short{"I"},
				Brief{"brief"},
				Usage{"usage"},
				Help{"help"},
				DefaultOn{},
				Terminates{},
				RunAfter{},
				func(Tri) int {
					return 0
				},
			},
			func(Tri) int {
				return 0
			},
		},
		{"node",
			Short{"n"},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Examples{"example1", "example2"},
			func(Tri) int {
				return 0
			},
		},
	},
}

// selftest is just a self test to make golint not tell me about unused things
func selftest() {
	brief := exampleTri[1].(Brief)
	fmt.Println(brief[0])
}
