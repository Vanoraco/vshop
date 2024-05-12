
const TOKEN_KEY = 'token';

export default {
  setToken(token) {
    localStorage.setItem(TOKEN_KEY, token);
  },
  getToken() {
    return localStorage.getItem(TOKEN_KEY);
  },
  removeToken() {
    localStorage.removeItem(TOKEN_KEY);
  },
  isAuthenticated() {
    const token = this.getToken();
    return token !== null && token !== undefined;
  }
};