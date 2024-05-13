const USER = 'user';

export default {
  setUser(user) {
    localStorage.setItem(USER, JSON.stringify(user));
  },
  getUser() {
    const userJSON = localStorage.getItem(USER);
    return userJSON ? JSON.parse(userJSON) : null;
  },
  removeUser() {
    localStorage.removeItem(USER);
  },
  isAuthenticated() {
    const user = this.getUser();
    return user !== null && user !== undefined;
  }
};
