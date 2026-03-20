package cli

type Flag string

const (
	FlagMaxArgs       Flag = "-n"
	FlagMaxLines      Flag = "-L"
	FlagNullDelimited Flag = "-0"
	FlagReplace       Flag = "-I"
	FlagPrint         Flag = "-t"
	FlagNoRunIfEmpty  Flag = "-r"
)

type FlagSpec struct {
	Name     Flag
	HasValue bool
}

var supportedFlags = map[Flag]FlagSpec{
	FlagMaxArgs:       {Name: FlagMaxArgs, HasValue: true},
	FlagMaxLines:      {Name: FlagMaxLines, HasValue: true},
	FlagNullDelimited: {Name: FlagNullDelimited, HasValue: false},
	FlagReplace:       {Name: FlagReplace, HasValue: true},
	FlagPrint:         {Name: FlagPrint, HasValue: false},
	FlagNoRunIfEmpty:  {Name: FlagNoRunIfEmpty, HasValue: false},
}

func isSupportedFlag(arg string) (FlagSpec, bool) {
	spec, ok := supportedFlags[Flag(arg)]
	return spec, ok
}
