package mstypes

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"math"
	"strings"
)

// RPCSID implements https://msdn.microsoft.com/en-us/library/cc230364.aspx
type RPCSID struct {
	Revision            uint8    // An 8-bit unsigned integer that specifies the revision level of the SID. This value MUST be set to 0x01.
	SubAuthorityCount   uint8    // An 8-bit unsigned integer that specifies the number of elements in the SubAuthority array. The maximum number of elements allowed is 15.
	IdentifierAuthority [6]byte  // An RPC_SID_IDENTIFIER_AUTHORITY structure that indicates the authority under which the SID was created. It describes the entity that created the SID. The Identifier Authority value {0,0,0,0,0,5} denotes SIDs created by the NT SID authority.
	SubAuthority        []uint32 `ndr:"conformant"` // A variable length array of unsigned 32-bit integers that uniquely identifies a principal relative to the IdentifierAuthority. Its length is determined by SubAuthorityCount.
}

// String returns the string representation of the RPC_SID.
func (s *RPCSID) String() string {
	var strb strings.Builder
	strb.WriteString("S-1-")

	b := append(make([]byte, 2, 2), s.IdentifierAuthority[:]...)
	// For a strange reason this is read big endian: https://msdn.microsoft.com/en-us/library/dd302645.aspx
	i := binary.BigEndian.Uint64(b)
	if i > math.MaxUint32 {
		fmt.Fprintf(&strb, "0x%s", hex.EncodeToString(s.IdentifierAuthority[:]))
	} else {
		fmt.Fprintf(&strb, "%d", i)
	}
	for _, sub := range s.SubAuthority {
		fmt.Fprintf(&strb, "-%d", sub)
	}
	return strb.String()
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"math"
	"strings"
>>>>>>> 4d7e5ad26 (update vendored files)
)

// RPCSID implements https://msdn.microsoft.com/en-us/library/cc230364.aspx
type RPCSID struct {
	Revision            uint8    // An 8-bit unsigned integer that specifies the revision level of the SID. This value MUST be set to 0x01.
	SubAuthorityCount   uint8    // An 8-bit unsigned integer that specifies the number of elements in the SubAuthority array. The maximum number of elements allowed is 15.
	IdentifierAuthority [6]byte  // An RPC_SID_IDENTIFIER_AUTHORITY structure that indicates the authority under which the SID was created. It describes the entity that created the SID. The Identifier Authority value {0,0,0,0,0,5} denotes SIDs created by the NT SID authority.
	SubAuthority        []uint32 `ndr:"conformant"` // A variable length array of unsigned 32-bit integers that uniquely identifies a principal relative to the IdentifierAuthority. Its length is determined by SubAuthorityCount.
}

// String returns the string representation of the RPC_SID.
func (s *RPCSID) String() string {
	var strb strings.Builder
	strb.WriteString("S-1-")

	b := append(make([]byte, 2, 2), s.IdentifierAuthority[:]...)
	// For a strange reason this is read big endian: https://msdn.microsoft.com/en-us/library/dd302645.aspx
	i := binary.BigEndian.Uint64(b)
	if i > math.MaxUint32 {
		fmt.Fprintf(&strb, "0x%s", hex.EncodeToString(s.IdentifierAuthority[:]))
	} else {
		fmt.Fprintf(&strb, "%d", i)
	}
	for _, sub := range s.SubAuthority {
		fmt.Fprintf(&strb, "-%d", sub)
	}
<<<<<<< HEAD
	return str
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	return str
=======
	return strb.String()
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
)

// RPCSID implements https://msdn.microsoft.com/en-us/library/cc230364.aspx
type RPCSID struct {
	Revision            uint8    // An 8-bit unsigned integer that specifies the revision level of the SID. This value MUST be set to 0x01.
	SubAuthorityCount   uint8    // An 8-bit unsigned integer that specifies the number of elements in the SubAuthority array. The maximum number of elements allowed is 15.
	IdentifierAuthority [6]byte  // An RPC_SID_IDENTIFIER_AUTHORITY structure that indicates the authority under which the SID was created. It describes the entity that created the SID. The Identifier Authority value {0,0,0,0,0,5} denotes SIDs created by the NT SID authority.
	SubAuthority        []uint32 `ndr:"conformant"` // A variable length array of unsigned 32-bit integers that uniquely identifies a principal relative to the IdentifierAuthority. Its length is determined by SubAuthorityCount.
}

// String returns the string representation of the RPC_SID.
func (s *RPCSID) String() string {
	var str string
	b := append(make([]byte, 2, 2), s.IdentifierAuthority[:]...)
	// For a strange reason this is read big endian: https://msdn.microsoft.com/en-us/library/dd302645.aspx
	i := binary.BigEndian.Uint64(b)
	if i >= 4294967296 {
		str = fmt.Sprintf("S-1-0x%s", hex.EncodeToString(s.IdentifierAuthority[:]))
	} else {
		str = fmt.Sprintf("S-1-%d", i)
	}
	for _, sub := range s.SubAuthority {
		str = fmt.Sprintf("%s-%d", str, sub)
	}
	return str
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
