<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Customer extends Model
{
    use HasFactory;
    protected $table = 'customer';

    protected $primaryKey = 'cid';

    public $timestamps = false;

    protected $fillable = [
        'name',
    ];

    protected $guarded = [
        'cid',
    ];

    protected $casts = [
        'name'  => 'string',
        'cid' => 'int',
    ];
}
