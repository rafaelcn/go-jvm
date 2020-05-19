package main

type (
	// ConstantPool ...
	ConstantPool struct {
	}

	// ClassFile represents the contents of a .class file
	ClassFile struct {
		Magic uint64
		Minor uint16
		Major uint16

        ConstantPoolCount uint16
        // The actual size of the ConstantPoolInfo is len(ConstantPoolCount)-1
        ConstantPoolInfo []uint16

		AccessFlags AFlags

        This *uint16
        Super *uint16

        InterfacesCount uint16
        // The actual size of interfaces is len(InterfacesCount)-1
        Interfaces []uint16

        FieldsCount uint16
        FieldsInfo []uint16

        MethodsCount uint16
        MethodsInfo []uint16
	}

	// AFlags is a type definition for access flags of a class
	AFlags uint16
)

const (
	// AccPublic ...
	AccPublic AFlags = 0x0001
	// AccFinal ...
	AccFinal AFlags = 0x0010
	// AccInterface implies that the defined file is an interface instead of a
	// class. If this flag is not defined then the given file is a class.
	AccInterface = 0x0200
	// AccAbstract ...
	AccAbstract AFlags = 0x0400
	// AccSynthetic ...
	AccSynthetic AFlags = 0x1000
	// AccAnnotation ...
	AccAnnotation AFlags = 0x2000
	// AccEnum ...
	AccEnum AFlags = 0x4000
)
