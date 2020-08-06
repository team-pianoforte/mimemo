export default async ({ store }) => {
  if (!store.getters['users/loggedIn']) {
    await store.dispatch('users/loginAsAnonymous')
  }
  await store.dispatch('users/prepare')
}
