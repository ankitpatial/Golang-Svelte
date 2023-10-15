import  { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
    schema: 'http://localhost:4000/query',
    documents: ['./src/lib/gql/*'],
    // emitLegacyCommonJSImports: false,
    ignoreNoDocuments: true, // for better experience with the watcher
    generates: {
        './src/lib/graph/': {
            preset: 'client',
            plugins: [ "urql-svelte-operations-store"],
            config: {
                withHooks: true,
            }
        },

    },
};

export default config;
