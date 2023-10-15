import {sveltekit} from '@sveltejs/kit/vite';
// import startInspector from 'svelte-inspector';
//
// startInspector({
//     editor: 'webstorm'
// });
const proxy = {};
[
    '/query',
    '/graphiql',
    '/me',
    '/login',
    '/logout',
    '/set-password',
    '/ws-msg',
    '/file-upload',
    '/file-download',
    '/PublicData'
].forEach((url) => {
    proxy[url] = {target: 'http://localhost:4000'};
});


export default {
    plugins: [sveltekit()],
    server: {
        proxy
    }
};
