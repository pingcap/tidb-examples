<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Order extends Model
{
    protected $table = 'order';

    protected $primaryKey = 'oid';

    public $timestamps = false;

    protected $fillable = [
        'cid',
        'price',
    ];

    protected $guarded = [
        'oid',
    ];

    protected $casts = [
        'uid'   => 'real',
        'price' => 'float',
    ];

    use HasFactory;
}
