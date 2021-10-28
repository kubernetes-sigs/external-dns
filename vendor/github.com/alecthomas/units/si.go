package units

// SI units.
type SI int64

// SI unit multiples.
const (
	Kilo SI = 1000
	Mega    = Kilo * 1000
	Giga    = Mega * 1000
	Tera    = Giga * 1000
	Peta    = Tera * 1000
	Exa     = Peta * 1000
)

func MakeUnitMap(suffix, shortSuffix string, scale int64) map[string]float64 {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	res := map[string]float64{
		shortSuffix: 1,
		// see below for "k" / "K"
		"M" + suffix: float64(scale * scale),
		"G" + suffix: float64(scale * scale * scale),
		"T" + suffix: float64(scale * scale * scale * scale),
		"P" + suffix: float64(scale * scale * scale * scale * scale),
		"E" + suffix: float64(scale * scale * scale * scale * scale * scale),
	}

	// Standard SI prefixes use lowercase "k" for kilo = 1000.
	// For compatibility, and to be fool-proof, we accept both "k" and "K" in metric mode.
	//
	// However, official binary prefixes are always capitalized - "KiB" -
	// and we specifically never parse "kB" as 1024B because:
	//
	// (1) people pedantic enough to use lowercase according to SI unlikely to abuse "k" to mean 1024 :-)
	//
	// (2) Use of capital K for 1024 was an informal tradition predating IEC prefixes:
	//     "The binary meaning of the kilobyte for 1024 bytes typically uses the symbol KB, with an
	//     uppercase letter K."
	//     -- https://en.wikipedia.org/wiki/Kilobyte#Base_2_(1024_bytes)
	//     "Capitalization of the letter K became the de facto standard for binary notation, although this
	//     could not be extended to higher powers, and use of the lowercase k did persist.[13][14][15]"
	//     -- https://en.wikipedia.org/wiki/Binary_prefix#History
	//     See also the extensive https://en.wikipedia.org/wiki/Timeline_of_binary_prefixes.
	if scale == 1024 {
		res["K"+suffix] = float64(scale)
	} else {
		res["k"+suffix] = float64(scale)
		res["K"+suffix] = float64(scale)
	}
	return res
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	return map[string]float64{
		shortSuffix:  1,
		"K" + suffix: float64(scale),
||||||| parent of 5ce8c7613 (update vendored files)
	return map[string]float64{
		shortSuffix:  1,
		"K" + suffix: float64(scale),
=======
	res := map[string]float64{
		shortSuffix: 1,
		// see below for "k" / "K"
>>>>>>> 5ce8c7613 (update vendored files)
		"M" + suffix: float64(scale * scale),
		"G" + suffix: float64(scale * scale * scale),
		"T" + suffix: float64(scale * scale * scale * scale),
		"P" + suffix: float64(scale * scale * scale * scale * scale),
		"E" + suffix: float64(scale * scale * scale * scale * scale * scale),
	}
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======

	// Standard SI prefixes use lowercase "k" for kilo = 1000.
	// For compatibility, and to be fool-proof, we accept both "k" and "K" in metric mode.
	//
	// However, official binary prefixes are always capitalized - "KiB" -
	// and we specifically never parse "kB" as 1024B because:
	//
	// (1) people pedantic enough to use lowercase according to SI unlikely to abuse "k" to mean 1024 :-)
	//
	// (2) Use of capital K for 1024 was an informal tradition predating IEC prefixes:
	//     "The binary meaning of the kilobyte for 1024 bytes typically uses the symbol KB, with an
	//     uppercase letter K."
	//     -- https://en.wikipedia.org/wiki/Kilobyte#Base_2_(1024_bytes)
	//     "Capitalization of the letter K became the de facto standard for binary notation, although this
	//     could not be extended to higher powers, and use of the lowercase k did persist.[13][14][15]"
	//     -- https://en.wikipedia.org/wiki/Binary_prefix#History
	//     See also the extensive https://en.wikipedia.org/wiki/Timeline_of_binary_prefixes.
	if scale == 1024 {
		res["K"+suffix] = float64(scale)
	} else {
		res["k"+suffix] = float64(scale)
		res["K"+suffix] = float64(scale)
	}
	return res
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	return map[string]float64{
		shortSuffix:  1,
		"K" + suffix: float64(scale),
||||||| parent of 6b7ce455e (update vendored files)
	return map[string]float64{
		shortSuffix:  1,
		"K" + suffix: float64(scale),
=======
	res := map[string]float64{
		shortSuffix: 1,
		// see below for "k" / "K"
>>>>>>> 6b7ce455e (update vendored files)
		"M" + suffix: float64(scale * scale),
		"G" + suffix: float64(scale * scale * scale),
		"T" + suffix: float64(scale * scale * scale * scale),
		"P" + suffix: float64(scale * scale * scale * scale * scale),
		"E" + suffix: float64(scale * scale * scale * scale * scale * scale),
	}
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======

	// Standard SI prefixes use lowercase "k" for kilo = 1000.
	// For compatibility, and to be fool-proof, we accept both "k" and "K" in metric mode.
	//
	// However, official binary prefixes are always capitalized - "KiB" -
	// and we specifically never parse "kB" as 1024B because:
	//
	// (1) people pedantic enough to use lowercase according to SI unlikely to abuse "k" to mean 1024 :-)
	//
	// (2) Use of capital K for 1024 was an informal tradition predating IEC prefixes:
	//     "The binary meaning of the kilobyte for 1024 bytes typically uses the symbol KB, with an
	//     uppercase letter K."
	//     -- https://en.wikipedia.org/wiki/Kilobyte#Base_2_(1024_bytes)
	//     "Capitalization of the letter K became the de facto standard for binary notation, although this
	//     could not be extended to higher powers, and use of the lowercase k did persist.[13][14][15]"
	//     -- https://en.wikipedia.org/wiki/Binary_prefix#History
	//     See also the extensive https://en.wikipedia.org/wiki/Timeline_of_binary_prefixes.
	if scale == 1024 {
		res["K"+suffix] = float64(scale)
	} else {
		res["k"+suffix] = float64(scale)
		res["K"+suffix] = float64(scale)
	}
	return res
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	return map[string]float64{
		shortSuffix:  1,
		"K" + suffix: float64(scale),
||||||| parent of 4d7e5ad26 (update vendored files)
	return map[string]float64{
		shortSuffix:  1,
		"K" + suffix: float64(scale),
=======
	res := map[string]float64{
		shortSuffix: 1,
		// see below for "k" / "K"
>>>>>>> 4d7e5ad26 (update vendored files)
		"M" + suffix: float64(scale * scale),
		"G" + suffix: float64(scale * scale * scale),
		"T" + suffix: float64(scale * scale * scale * scale),
		"P" + suffix: float64(scale * scale * scale * scale * scale),
		"E" + suffix: float64(scale * scale * scale * scale * scale * scale),
	}
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======

	// Standard SI prefixes use lowercase "k" for kilo = 1000.
	// For compatibility, and to be fool-proof, we accept both "k" and "K" in metric mode.
	//
	// However, official binary prefixes are always capitalized - "KiB" -
	// and we specifically never parse "kB" as 1024B because:
	//
	// (1) people pedantic enough to use lowercase according to SI unlikely to abuse "k" to mean 1024 :-)
	//
	// (2) Use of capital K for 1024 was an informal tradition predating IEC prefixes:
	//     "The binary meaning of the kilobyte for 1024 bytes typically uses the symbol KB, with an
	//     uppercase letter K."
	//     -- https://en.wikipedia.org/wiki/Kilobyte#Base_2_(1024_bytes)
	//     "Capitalization of the letter K became the de facto standard for binary notation, although this
	//     could not be extended to higher powers, and use of the lowercase k did persist.[13][14][15]"
	//     -- https://en.wikipedia.org/wiki/Binary_prefix#History
	//     See also the extensive https://en.wikipedia.org/wiki/Timeline_of_binary_prefixes.
	if scale == 1024 {
		res["K"+suffix] = float64(scale)
	} else {
		res["k"+suffix] = float64(scale)
		res["K"+suffix] = float64(scale)
	}
	return res
>>>>>>> 4d7e5ad26 (update vendored files)
}
