def _gcs_upload_impl(ctx):
  targets = []
  for target in ctx.files.data:
    targets.append(target.short_path)

  ctx.file_action(
      output = ctx.outputs.targets,
      content  = "\n".join(targets),
  )

  ctx.file_action(
      content = "%s --manifest %s --root $PWD -- $@" % (
          ctx.attr.uploader.files_to_run.executable.short_path,
          ctx.outputs.targets.short_path,
      ),
      output = ctx.outputs.executable,
      executable = True,
  )

  return struct(
      runfiles = ctx.runfiles(
          files = ctx.files.data + ctx.files.uploader +
            [ctx.version_file, ctx.outputs.targets]
          )
      )

gcs_upload = rule(
    attrs = {
        "data": attr.label_list(
            mandatory = True,
            allow_files = True,
        ),
        "uploader": attr.label(
            default = Label("//defs:gcs_uploader"),
            allow_files = True,
        ),
    },
    executable = True,
    outputs = {
        "targets": "%{name}-targets.txt",
    },
    implementation = _gcs_upload_impl,
)
