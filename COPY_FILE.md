# File Copy Feature

This library now supports copying files within the same bucket or across different buckets, matching the functionality of the official Supabase JavaScript client.

## Features

- ✅ Copy files within the same bucket
- ✅ Copy files across different buckets
- ✅ Optional metadata copying control
- ✅ Support for files up to 5 GB

## API Reference

### CopyFile Method

```go
func (c *Client) CopyFile(
    bucketId string,
    sourceKey string,
    destinationKey string,
    options ...CopyFileOptions,
) (FileCopyResponse, error)
```

#### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `bucketId` | string | Yes | The source bucket ID |
| `sourceKey` | string | Yes | The source file path, including the file name (e.g., `folder/subfolder/filename.png`) |
| `destinationKey` | string | Yes | The destination file path, including the file name (e.g., `folder/subfolder/copy-filename.png`) |
| `options` | CopyFileOptions | No | Optional parameters for cross-bucket copying and metadata handling |

#### CopyFileOptions

```go
type CopyFileOptions struct {
    // DestinationBucket is the target bucket id. If not specified, defaults to source bucket
    DestinationBucket *string
    // CopyMetadata determines whether to copy metadata from source. Defaults to true.
    CopyMetadata *bool
}
```

#### Response

```go
type FileCopyResponse struct {
    Key     string `json:"Key"`           // The destination file path
    Message string `json:"message,omitempty"` // Optional message
    Error   string `json:"error,omitempty"`   // Error message if any
}
```

## Usage Examples

### Example 1: Copy file within the same bucket

```go
package main

import (
    "fmt"
    "log"

    storage_go "github.com/supabase-community/storage-go"
)

func main() {
    client := storage_go.NewClient("https://your-project.supabase.co/storage/v1", "your-api-key", nil)

    response, err := client.CopyFile(
        "avatars",                   // source bucket
        "public/avatar1.png",        // source file
        "private/avatar1-copy.png",  // destination file
    )

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("File copied: %s\n", response.Key)
}
```

### Example 2: Copy file across buckets

```go
destinationBucket := "avatars-backup"

response, err := client.CopyFile(
    "avatars",                // source bucket
    "public/avatar1.png",     // source file
    "backup/avatar1.png",     // destination file
    storage_go.CopyFileOptions{
        DestinationBucket: &destinationBucket,
    },
)

if err != nil {
    log.Fatal(err)
}

fmt.Printf("File copied across buckets: %s\n", response.Key)
```

### Example 3: Copy file without metadata

```go
copyMetadata := false

response, err := client.CopyFile(
    "avatars",                    // source bucket
    "public/avatar1.png",         // source file
    "public/avatar1-no-meta.png", // destination file
    storage_go.CopyFileOptions{
        CopyMetadata: &copyMetadata,
    },
)

if err != nil {
    log.Fatal(err)
}

fmt.Printf("File copied without metadata: %s\n", response.Key)
```

### Example 4: Copy file across buckets without metadata

```go
destinationBucket := "avatars-backup"
copyMetadata := false

response, err := client.CopyFile(
    "avatars",                    // source bucket
    "public/avatar1.png",         // source file
    "backup/avatar1-no-meta.png", // destination file
    storage_go.CopyFileOptions{
        DestinationBucket: &destinationBucket,
        CopyMetadata:      &copyMetadata,
    },
)

if err != nil {
    log.Fatal(err)
}

fmt.Printf("File copied: %s\n", response.Key)
```

## Permissions Required

For a user to copy objects, they need:
- `select` permission on the source object
- `insert` permission on the destination object

Example RLS policies:

```sql
-- Allow users to read their own objects
CREATE POLICY "User can select their own objects"
ON storage.objects
FOR SELECT
TO authenticated
USING (owner_id = auth.uid());

-- Allow users to upload in their own folders
CREATE POLICY "User can upload in their own folders"
ON storage.objects
FOR INSERT
TO authenticated
WITH CHECK ((storage.folder(name))[1] = auth.uid());
```

## Limitations

- Maximum file size: 5 GB
- Requires proper RLS policies for both source and destination
- When copying across buckets, both buckets must be accessible

## API Endpoint

The method calls the following Supabase Storage API endpoint:

```
POST /object/copy
```

Request body:
```json
{
  "bucketId": "source-bucket",
  "sourceKey": "path/to/source.png",
  "destinationKey": "path/to/destination.png",
  "destinationBucket": "optional-destination-bucket",
  "copyMetadata": true
}
```

## Comparison with JavaScript Client

This implementation matches the behavior of the official Supabase JavaScript client:

**JavaScript:**
```javascript
await supabase.storage
  .from('avatars')
  .copy('public/avatar1.png', 'private/avatar2.png', {
    destinationBucket: 'avatars2',
  })
```

**Go:**
```go
destinationBucket := "avatars2"
client.CopyFile("avatars", "public/avatar1.png", "private/avatar2.png",
    storage_go.CopyFileOptions{
        DestinationBucket: &destinationBucket,
    })
```
