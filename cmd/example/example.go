package main

import (
	. "git.parallelcoin.io/tri"
)

type exampleConf struct {
	datadir string
}

var cfg = exampleConf{
	datadir: "/not/the/default/path",
}

var cfg2 = exampleConf{
	datadir: "/not/the/default/path",
}

// dot import used on tri as it should be for declarations, or every damned name has to have tri. in front of it. This is ok in other places, but not here, as you can obviously see:
var exampleTri = Tri{
	"appname",
	Brief{"brief"},
	Version{0, 1, 1, "alpha"},
	DefaultCommand{"ctl"},
	Var{"datadir",
		Short{'d'},
		Brief{"brief"},
		Usage{"usage"},
		Help{"help"},
		Default{"~/.pod"},
		Slot{&cfg.datadir},
	},
	Trigger{"init",
		Short{'I'},
		Brief{"brief"},
		Usage{"usage"},
		Help{"help"},
		DefaultOn{},
		RunAfter{},
		func(*Tri) int {
			return 0
		},
	},
	Commands{
		{"ctl",
			Short{'c'},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Examples{
				"example 1", "explaining text",
				"example 2", "explaining text",
			},
			Var{"datadir",
				Short{'d'},
				Brief{"brief"},
				Usage{"usage"},
				Help{"help"},
				Default{"~/.pod"},
				Group{"groupname"},
				Slot{&cfg.datadir, &cfg2.datadir},
			},
			Trigger{"wallet",
				Short{'w'},
				Brief{"brief"},
				Usage{"usage"},
				Help{"help"},
				DefaultOn{},
				Terminates{},
				func(*Tri) int {
					return 0
				},
			},
			func(*Tri) int {
				return 0
			},
		},
		{"node",
			Short{'n'},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Examples{"example1", "example2"},
			func(*Tri) int {
				return 0
			},
		},
	},
}
