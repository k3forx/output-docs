package main

type BasicInterface interface {
	String() string
}

type BasicInterfaceImpl struct{}

func (i BasicInterfaceImpl) String() string {
	return "via basic interface!"
}

type GeneralInterface interface {
	String() string
	GeneralInterfaceImpl
}

type GeneralInterfaceImpl struct{}

func (i GeneralInterfaceImpl) String() string {
	return "via general interface!"
}

func main() {
	bi := BasicInterfaceImpl{}
	StringFromBasicInterface(bi)

	gi := GeneralInterfaceImpl{}
	StringFromGeneralInterface[GeneralInterfaceImpl](gi)
}

func StringFromBasicInterface(bi BasicInterface) {
	_ = bi.String()
}

func StringFromGeneralInterface[GI GeneralInterface](gi GI) {
	_ = gi.String()
}
