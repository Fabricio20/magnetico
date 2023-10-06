# API

This fork implements a ZeroMQ interface that acts as a PUB. It constantly publishes messages in the following format
to subscribers:

```json
{
  "infoHash": "f84b51f0d2c3455ab5dabb6643b4340234cd036e",
  "name": "Big_Buck_Bunny_1080p_surround_frostclick.com_frostwire.com",
  "files": [
    {
      "size": 928670754,
      "path": "Big_Buck_Bunny_1080p_surround_FrostWire.com.avi"
    },
    {
      "size": 5008,
      "path": "PROMOTE_YOUR_CONTENT_ON_FROSTWIRE_01_06_09.txt"
    },
    {
      "size": 3456234,
      "path": "Pressrelease_BickBuckBunny_premiere.pdf"
    },
    {
      "size": 180,
      "path": "license.txt"
    }
  ]
}
```

By default, the ZeroMQ PUB binds to port `2222` on `UDP`. It is not currently possible to configure this, however
you can always remap using docker.
