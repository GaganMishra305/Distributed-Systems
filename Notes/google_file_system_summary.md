# Google File System (GFS) Architecture
- **Single Master**: Holds all file metadata in memory (fast lookups).
- **Chunkservers**: Store 64MB chunks replicated 3x across racks.
- **Client Direct I/O**: Client asks master for chunk location, then streams data directly to/from chunkservers.
- **Record Append**: Guaranteed atomic append even under concurrent appenders.
