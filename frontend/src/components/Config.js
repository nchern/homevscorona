class Config {

  getApiHost() {
      if ( process.env.NODE_ENV === 'development' ) {
          return 'http://localhost:8080';
      }
      return 'http://homevscorona.us.to';
  }
}

export default new Config();
