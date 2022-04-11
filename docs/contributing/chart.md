# Helm Chart

## Chart Changes

When contributing chart changes please follow the same process as when contributing other content but also please **DON'T** modify _Chart.yaml_ in the PR as this would result in a chart release when merged and will mean that your PR will need modifying before it can be accepted. The chart version will be updated as part of the PR to release the chart.

Please **DO** add your changes to the _CHANGELOG.md_ file in the chart directory under the `## [UNRELEASED]` section, if there isn't an uncommented `## [UNRELEASED]` section please copy the commented out template and use that.
