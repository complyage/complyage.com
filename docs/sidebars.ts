// sidebars.ts

import type { SidebarsConfig } from '@docusaurus/plugin-content-docs';

const sidebars: SidebarsConfig = {
  docs: [
    {
      type: 'category',
      label: 'Client',
      collapsed: false,
      items: [
        'client/overview',
        'client/installation',
        'client/configuration',
        'client/usage',
        {
          type: 'category',
          label: 'Components',
          items: [
            'client/components/auth',
            'client/components/forms',
            'client/components/errors'
          ],
        },
        'client/faq'
      ],
    },
    {
      type: 'category',
      label: 'OAuth',
      collapsed: false,
      items: [
        'oauth/overview',
        'oauth/authorization-code-flow',
        'oauth/client-example',
        'oauth/scopes',
        'oauth/errors',
        'oauth/faq'
      ],
    },
    {
      type: 'category',
      label: 'Agent',
      collapsed: false,
      items: [
        'agent/overview',
        'agent/installation',
        'agent/verification-steps',
        'agent/integration',
        'agent/troubleshooting'
      ],
    },
    {
      type: 'category',
      label: 'End User',
      collapsed: false,
      items: [
        'enduser/overview',
        'enduser/setup',
        'enduser/account',
        'enduser/verification',
        'enduser/faq'
      ],
    }
  ],
};

export default sidebars;
