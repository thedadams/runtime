/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  "sidebar": [
    "home",
    {
      "type": "category",
      "label": "Installation",
      "items": [
        "installation/installing",
        "installation/options",
        "installation/upgrading",
        "installation/uninstalling"
      ],
      "collapsed": true
    },
    "getting-started",
    {
      "type": "category",
      "label": "Administration",
      "items": [
        "admin/volumeclasses",
        "admin/computeclasses",
        "admin/alpha-image-allow-rules",
      ]
    },
    {
      "type": "category",
      "label": "Architecture",
      "items": [
        "architecture/ten-thousand-foot-view",
        "architecture/security-considerations"
      ],
      "collapsed": true
    },
    {
      "type": "category",
      "label": "Reference",
      "items": [
        {
          "type": "category",
          "label": "Command Line",
          "items": [
            "reference/command-line/acorn",
            "reference/command-line/acorn_all",
            "reference/command-line/acorn_build",
            "reference/command-line/acorn_check",
            "reference/command-line/acorn_container",
            "reference/command-line/acorn_container_kill",
            "reference/command-line/acorn_credential",
            "reference/command-line/acorn_credential_login",
            "reference/command-line/acorn_credential_logout",
            "reference/command-line/acorn_events",
            "reference/command-line/acorn_exec",
            "reference/command-line/acorn_image",
            "reference/command-line/acorn_image_rm",
            "reference/command-line/acorn_info",
            "reference/command-line/acorn_install",
            "reference/command-line/acorn_login",
            "reference/command-line/acorn_logout",
            "reference/command-line/acorn_logs",
            "reference/command-line/acorn_project",
            "reference/command-line/acorn_project_create",
            "reference/command-line/acorn_project_rm",
            "reference/command-line/acorn_project_use",
            "reference/command-line/acorn_ps",
            "reference/command-line/acorn_pull",
            "reference/command-line/acorn_push",
            "reference/command-line/acorn_offerings",
            "reference/command-line/acorn_offerings_computeclasses",
            "reference/command-line/acorn_offerings_volumeclasses",
            "reference/command-line/acorn_render",
            "reference/command-line/acorn_rm",
            "reference/command-line/acorn_run",
            "reference/command-line/acorn_secret",
            "reference/command-line/acorn_secret_create",
            "reference/command-line/acorn_secret_encrypt",
            "reference/command-line/acorn_secret_reveal",
            "reference/command-line/acorn_secret_rm",
            "reference/command-line/acorn_start",
            "reference/command-line/acorn_stop",
            "reference/command-line/acorn_tag",
            "reference/command-line/acorn_uninstall",
            "reference/command-line/acorn_update",
            "reference/command-line/acorn_volume",
            "reference/command-line/acorn_volume_rm",
            "reference/command-line/acorn_wait"
          ]
        },
      ],
      "collapsed": true
    },
    "faq"
  ]

};

module.exports = sidebars;
