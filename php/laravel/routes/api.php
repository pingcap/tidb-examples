<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\customerController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::get('/customer/{id}', 'App\Http\Controllers\CustomerController@getByCid');
Route::post('/customer', 'App\Http\Controllers\CustomerController@insert');

Route::post('/order', 'App\Http\Controllers\OrderController@insert');
Route::delete('/order/{oid}', 'App\Http\Controllers\OrderController@delete');
Route::post('/order/{oid}','App\Http\Controllers\OrderController@updateByOid');
Route::get('/order','App\Http\Controllers\OrderController@queryByCid');
