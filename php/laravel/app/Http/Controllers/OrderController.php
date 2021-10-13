<?php

namespace App\Http\Controllers;

use App\Models\Order;
use Illuminate\Http\Request;

class OrderController extends Controller
{
    public function insert(Request $request) {
        return Order::create([
            'cid' => $request->cid,
            'price' => $request->price,
        ]);
    }

    public function delete($cid) {
        return Order::withTrashed()
            ->where('cid', $cid)
            ->get();
    }

    public function updateByOid(Request $request) {
        Order::where('oid', $request->oid)->update(['price' => $request->price]);   
    }

    public function queryByCid(Request $request) {
        return Order::join('customer', 'order.cid', '=', 'customer.cid')->where('order.cid', $request->cid)->get();
    }
}
