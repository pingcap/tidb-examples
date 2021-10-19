<?php

namespace App\Http\Controllers;

use App\Models\Customer;
use Illuminate\Http\Request;

class CustomerController extends Controller
{
    public function getByCid($cid)
    {
        $customer_info = Customer::where('cid',$cid)->get();
        if ($customer_info->count() > 0){
            return $customer_info;
        }
        return abort(404);
    }

    public function insert(Request $request) {
        return Customer::create(['name' => $request->name]);
    }
}
