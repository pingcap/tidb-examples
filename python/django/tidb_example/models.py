from django.db import models


class Orders(models.Model):
    oid = models.AutoField(primary_key=True)
    uid = models.BigIntegerField(null=False)
    price = models.FloatField()

class GENDER(models.TextChoices):
    FEMALE = 'Female'
    MALE = 'Male'

class Users(models.Model):
    uid = models.AutoField(primary_key=True)
    name = models.CharField(max_length=250)
    gender = models.CharField(choices=GENDER.choices, max_length=6,default=GENDER.FEMALE)
