module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: './tsconfig.eslint.json',
    tsconfigRootDir: __dirname,
  },
  plugins: ['@typescript-eslint', 'import'],
  extends: [
    'plugin:prettier/recommended',
    'plugin:react/jsx-runtime',
    'prettier',
  ],
  env: {
    jest: true,
    browser: true,
  },
  globals: {
    __ENV: true,
  },
  rules: {
    'import/no-cycle': 'error',
    'prettier/prettier': 'error',
    'no-underscore-dangle': 'off',
    'no-plusplus': 'off',
    'import/no-named-as-default-member': 'warn',
    '@typescript-eslint/ban-ts-comment': 'warn',
    '@typescript-eslint/interface-name-prefix': 0,
    '@typescript-eslint/no-redeclare': ['error', { builtinGlobals: false }],
    '@typescript-eslint/default-param-last': 'warn',
    'no-template-curly-in-string': 'off',
    'import/order': [
      'error',
      {
        'newlines-between': 'always',
        groups: [
          ['builtin', 'unknown'],
          'external',
          'internal',
          ['parent', 'sibling', 'index'],
        ],

        pathGroupsExcludedImportTypes: [],
        pathGroups: [],
      },
    ],
  },
  settings: {
    react: {
      version: 'detect',
    },
    'import/resolver': {
      typescript: {},
    },
  },
  overrides: [
    {
      files: ['*.ts', '*.tsx'],
      extends: ['airbnb-typescript', 'plugin:prettier/recommended'],
      rules: {
        'import/no-extraneous-dependencies': [
          'error',
          {
            devDependencies: true,
          },
        ],
      },
    },
    {
      files: ['*.jsx', '*.tsx'],
      plugins: ['react'],
      rules: {
        'react/require-default-props': 'off',
        'react/jsx-props-no-spreading': 'off',
        'react/static-property-placement': [
          'error',
          'property assignment',
          {
            defaultProps: 'static public field',
          },
        ],
      },
    },
    {
      files: ['*.tsx'],
      rules: {
        'react/prop-types': 'off',
      },
    },
    {
      files: ['*.mdx'],
      plugins: ['unused-imports', 'react'],
      extends: 'plugin:mdx/recommended',
      rules: {
        'unused-imports/no-unused-imports': 'error',
        'react/jsx-uses-react': 'off',
        'react/jsx-uses-vars': 'error',
      },
      settings: {
        'mdx/code-blocks': true,
        'mdx/language-mapper': {},
      },
    },
  ],
  ignorePatterns: ['.eslintrc.js'],
};
