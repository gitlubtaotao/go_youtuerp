const path = require('path');

const HtmlWebpackPlugin = require('html-webpack-plugin'); //自动生成html文件

//A webpack plugin to remove your build folder(s) before building
// https://github.com/johnagan/clean-webpack-plugin
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const webpack = require('webpack');
module.exports = {
    entry: {
        main: './src/app.js',
    },
    module: {
        rules: [
            // {
            // 	test: /\.(jpg|png|gif)$/,
            // 	use:{
            // 		loader: 'file-loader',
            // 		options:{
            // 			//placeholder 占位符
            // 			name: '[name]_[hash].[ext]',
            // 			outputPath: 'images/'
            // 		}
            // 	}
            // },
            {
                test: /\.(eot|ttf|svg|woff)$/,
                use: {
                    loader: 'file-loader'
                }
            },
            {
                test: /\.(jpg|png|gif)$/,
                use: {
                    loader: 'url-loader',
                    options: {
                        //placeholder 占位符
                        name: '[name]_[hash].[ext]',
                        outputPath: 'images/',
                        limit: 2048,
                    }
                }
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader', 'postcss-loader']
            },
            {
                test: /\.scss$/,
                use: ['style-loader',
                    {
                        loader: 'css-loader',
                        options: {
                            importLoaders: 2, //
                            modules: true //css 模块化打包
                        }
                    },
                    'sass-loader', 'postcss-loader']
            },
            {
                test: /\.vue$/,
                use: {
                    loader: 'vue-loader'
                }
            },
            //babel-loader: es6 to es5
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: "babel-loader",
                // options:{

                // 	// @babel/polyfill
                // 	// "presets": [["@babel/preset-env",
                // 	// 	{
                // 	// 	useBuiltIns: "usage",
                // 	// 	targets:{
                // 	// 		chrome: "67"
                // 	// 	}}
                // 	// ]]
                // 	//A plugin that enables the re-use of Babel's injected helper code to save on codesize

                // 	"plugins": [["@babel/plugin-transform-runtime",{
                // 		"absoluteRuntime": false,
                //    	"corejs": 2,
                //     	"helpers": true,
                //      "regenerator": true,
                //      "useESModules": false,
                //      "version": "7.0.0-beta.0"
                // 	}]]

                // }
            }]
    },

    plugins: [
        new CleanWebpackPlugin({}), //清空上次打包的文件
        // new HtmlWebpackPlugin({
        //     template: 'src/index.html', //指定生成html文件饿模版
        //     title: 'plugins'
        // }),
    ],
    optimization: {
        splitChunks: {
            chunks: 'all'
        }
    },
    output: {
        //filename: 'bundle.js'，
        filename: '[name].js',
        // path: path.resolve(__dirname, '../dist'),
        // //publicPath: 'http://cdn-webpacker.com', //生成url
        // publicPath: '/assets/' //生成一个新文件夹

    }
};