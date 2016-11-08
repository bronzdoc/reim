# Reim - image resizer

## Install
```go get github.com/bronzdoc/reim```

## Usage

```$ reim image.jpg ```

or passing flags

``` $ reim -width 200 -height 200 image.jpg```

#### flags
```
  -height uint
        New image height (default 128)
  -out string
        Generated image name
  -width uint
        New image width (default 128)
```

The default image size is 128x128 because reim was thought on slack emojis :trollface:

