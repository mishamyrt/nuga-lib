# Backlight custom effect

In addition to the preset effects, the keyboard is able to display a custom one. It is static (has no animation) and allows you to configure separate colors for each key.

## Read

To read a custom effect, you have to send a request:

```sh
0x05 0x89 0xBC 0x00 0x00 0x00
```

The response will be a sequence of 1050 bytes. The colors for the keys are encoded in these bytes.

The first will be the header, which is always `0x89`. Followed by 6 zeros.

Then comes the effect colors. Keys positions are sampled from byte 7. You can see an example of positions in [template_backlight.go](../layout/template_backlight.go).

Colors are written in a rather strange way: first comes red, after 126 bytes comes green and after another 126 bytes comes blue.
That is, if you have white color written for the first and second button, then 255 will be written to the bytes under the indexes:

```sh
# First
7, 133, 259
# Second
8, 134, 260
```

## Write

To record the effect you need to send a sequence of bytes in the same format but with a different header. Cut off the first 7 bytes and add the command:

```sh
0x06 0x09 0xBC 0x00 0x40 0x00 0x00 0x00 # command
0xFF 0xFF 0x00 0x00 0x00 0x00 0x00 0x00 # payload
...
```

## Pages

It seems that the keyboard can store several pages of custom effects, but I haven't figured out how to activate them yet. There are the following commands to work with them:

```sh
# Read
0x05 0x89 0xC0 0x00 0x00 0x00
# Write
0x06 0x09 0xC0 0x00 0x40 0x00 0x00 0x00
```