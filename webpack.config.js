module.exports = {
  entry: './public/js/index.js',
  output: {
    path: __dirname,
    filename: './build/public/js/bundle.js'
  },

  module: {
    loaders: [
      {
         test: /\.jsx?$/,
         loader: 'babel-loader',
         exclude: /node_modules/,
         query  :{
           presets:['react','es2015']
         }
      }
    ]
  }
}
