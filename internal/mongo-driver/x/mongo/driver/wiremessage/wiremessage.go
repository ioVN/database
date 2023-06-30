package wiremessage

const (
	// DefaultZlibLevel is the default level for zlib compression
	DefaultZlibLevel = 6
	// DefaultZstdLevel is the default level for zstd compression.
	// Matches https://github.com/wiredtiger/wiredtiger/blob/f08bc4b18612ef95a39b12166abcccf207f91596/ext/compressors/zstd/zstd_compress.c#L299
	DefaultZstdLevel = 6
)
