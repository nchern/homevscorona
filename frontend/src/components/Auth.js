class Auth {
  constructor() {
      this.token = localStorage.getItem('ggToken');
      if (this.token == null) {
          this.authenticated = false;
      } else {
          this.authenticated = true;
      }
  }
  isAuthenticated() {
    return this.authenticated;
  }
}

export default new Auth();