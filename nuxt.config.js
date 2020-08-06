import colors from 'vuetify/es5/util/colors'

export default {
  mode: 'spa',
  target: 'static',

  server: {
    host: '0.0.0.0',
    port: 8000,
  },
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
    ],
    script: [{ src: '/ncmb.min.js' }],
  },
  env: {
    ncmb: {
      applicationKey: process.env.NCMB_APPLICATION_KEY,
      clientKey: process.env.NCMB_CLIENT_KEY,
    },
  },
  plugins: [
  ],
  srcDir: 'src',
  components: true,
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    '@nuxtjs/eslint-module',
    '@nuxtjs/vuetify',
  ],
  modules: [
    '@nuxtjs/pwa',
  ],
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: true,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3,
        },
      },
    },
  },
  /*
  ** Build configuration
  ** See https://nuxtjs.org/api/configuration-build/
  */
  build: {
  },
}
