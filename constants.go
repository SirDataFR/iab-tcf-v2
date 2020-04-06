package iabtcfv2

type SegmentType int

const (
	SegmentTypeUndefined        SegmentType = -1
	SegmentTypeCoreString       SegmentType = 0
	SegmentTypeDisclosedVendors SegmentType = 1
	SegmentTypeAllowedVendors   SegmentType = 2
	SegmentTypePublisherTC      SegmentType = 3
)

type TcfVersion int

const (
	TcfVersionUndefined TcfVersion = -1
	TcfVersion1         TcfVersion = 1
	TcfVersion2         TcfVersion = 2
)
