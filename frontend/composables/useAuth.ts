import { ref } from 'vue';

const loggedIn = ref(true);

export function useAuth() {
  return { loggedIn };
}

export function setLoggedIn(status: boolean) {
  loggedIn.value = status;
  return loggedIn.value;
}