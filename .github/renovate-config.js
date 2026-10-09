"use strict";
// https://github.com/renovatebot/github-action/blob/main/.github/renovate.json
// https://docs.renovatebot.com/configuration-options/

module.exports = {
  "extends": [":disableRateLimiting", ":semanticCommits"],
  "assigneesFromCodeOwners": true,
  "gitAuthor": "Renovate Bot <bot@external-dns.com>",
  "onboarding": false,
  "platform": "github",
  "repositories": [
    "kubernetes-sigs/external-dns"
  ],
  "printConfig": false,
  "prConcurrentLimit": 0,
  "prHourlyLimit": 0,
  "minimumReleaseAge": "3 days",
  "pruneStaleBranches": true,
  "dependencyDashboard": false,
  "requireConfig": false,
  "rebaseWhen": "behind-base-branch",
  "baseBranches": ["master"],
  "recreateWhen": "always",
  "semanticCommits": "enabled",
  "labels": ["{{depType}}", "datasource::{{datasource}}", "type::{{updateType}}", "manager::{{manager}}"], // can be overridden per packageRule
  "addLabels": ["renovate-bot"], // cannot be overridden, any packageRule config extends this
  "packageRules": [
  ],
  "enabledManagers": [ // supported managers https://docs.renovatebot.com/modules/manager/
    "mise", // reads mise.toml
    "custom.regex"
  ],
  // 'mise lock' refreshes mise.lock; Renovate blocks it without this opt-in
  "allowedUnsafeExecutions": ["mise", "gradleWrapper"],
  "customManagers": [ // https://docs.renovatebot.com/modules/manager/regex/
    {
      // to capture registry.k8s.io/external-dns/external-dns:<version> in *.md files
      "customType": "regex",
      "managerFilePatterns": [
        "/.*\\.md$/"
      ],
      "matchStrings": [
        "(?<depName>registry.k8s.io\/external-dns\/external-dns):(?<currentValue>.*)"
      ],
      "depNameTemplate": "kubernetes-sigs/external-dns",
      "datasourceTemplate": "github-releases",
      "versioningTemplate": "semver"
    },
    {
      "customType": "regex",
      "managerFilePatterns": ["/.*/"],
      "matchStrings": [
        "datasource=(?<datasource>.*?) depName=(?<depName>.*?)( versioning=(?<versioning>.*?))?\\s.*?_VERSION=(?<currentValue>.*)\\s"
      ],
      "versioningTemplate": "{{#if versioning}}{{{versioning}}}{{else}}semver{{/if}}",
    },
    {
      // mise CLI version in the dev-env composite action
      "customType": "regex",
      "managerFilePatterns": ["/^\\.github/actions/.+/action\\.yml$/"],
      "matchStrings": [
        "# renovate: datasource=(?<datasource>\\S+) depName=(?<depName>\\S+)\\s+default: '(?<currentValue>[^']+)'"
      ],
      "versioningTemplate": "loose",
      "extractVersionTemplate": "^v?(?<version>.*)$"
    },
  ]
};
