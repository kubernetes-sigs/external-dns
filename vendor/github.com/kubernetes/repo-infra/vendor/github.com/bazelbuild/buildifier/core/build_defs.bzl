_GO_TOOL = "@io_bazel_rules_go//go/toolchain:go_tool"

def go_yacc(src, out, visibility=None):
  native.genrule(
    name = src + ".go_yacc",
    srcs = [src],
    outs = [out],
    tools = [_GO_TOOL],
    cmd = ("export GOROOT=$$(dirname $(location " + _GO_TOOL + "))/..;" +
           " $(location " + _GO_TOOL + ") tool yacc " +
           " -o $(location " + out + ") $(SRCS)"),
    visibility = visibility,
    local = 1,
   )
