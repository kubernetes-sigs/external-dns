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
  "stabilityDays": 3,
  "pruneStaleBranches": true,
  "recreateClosed": true,
  "dependencyDashboard": false,
  "requireConfig": false,
  "rebaseWhen": "behind-base-branch",
  "baseBranches": ["master", "main"],
  "recreateWhen": "always",
  "semanticCommits": "enabled",
  "pre-commit": {
    "enabled": true
  },
  "labels": ["{{depType}}", "datasource::{{datasource}}", "type::{{updateType}}", "manager::{{manager}}"], // can be overridden per packageRule
  "addLabels": ["renovate-bot"], // cannot be overridden, any packageRule config extends this
  "packageRules": [
    {
      "groupName": "pre-commit",
      "matchManagers": ["pre-commit"],
      "addLabels": ["pre-commit", "skip-release"]
    },
  ],
  "enabledManagers": [ // supported managers https://docs.renovatebot.com/modules/manager/
    "regex",
    "pre-commit"
  ],
  "customManagers": [ // https://docs.renovatebot.com/modules/manager/regex/
    {
      // to capture registry.k8s.io/external-dns/external-dns:<version> in *.md files
      "customType": "regex",
      "fileMatch": [
        ".*\\.md$"
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
      "fileMatch": [".*"],
      "matchStrings": [
        "datasource=(?<datasource>.*?) depName=(?<depName>.*?)( versioning=(?<versioning>.*?))?\\s.*?_VERSION=(?<currentValue>.*)\\s"
      ],
      "versioningTemplate": "{{#if versioning}}{{{versioning}}}{{else}}semver{{/if}}",
    },
  ]
};
