

//sourceMap: 映射关系 最佳实践
//{mode: 'production',devtool: 'cheap-module-source-map'}
//{mode: 'development',devtool: 'cheap-module-eval-source-map'}
const webpack = require('webpack');

const merge = require('webpack-merge');
const commonConfig = require('./webpack.common.js');

const devConfig = {
	mode:'development',
	// devtool: 'none',//关闭开发模式 sourceMap
	devtool:  'cheap-module-eval-source-map',
 	//webpack-dev-server
 	devServer:{
 		contentBase: './dist',
 		open: false, //自动打开游览器
 		port: 8080, 		
 		proxy: {
 			'/api': "http: //localhost: 3000"
 		},
 		
 		//热模块更新
 		hot: true,
 		hotOnly: true
 	},
   	plugins: [
	   	//热模块更新
	   	new webpack.HotModuleReplacementPlugin(),
   	 
   	],

   	// Tree Shaking
   	optimization:{
   		usedExports: true 
   	}
}

module.exports = merge(commonConfig,devConfig)
