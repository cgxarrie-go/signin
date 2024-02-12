# signin

CLI for reservation of desks via signin app API
(https://companion.signin.app/)


## commands

 - [book](#book)
 - [cancel booking](#cancel-booking)
 - [list free spaces](#list-free-spaces-for-a-given-date)
 - [list bookings](#list-bookings-up-to-a-given-date)
 - [config](#config)
    - [bearer](#config-bearer)

### book
Book a space for a given date

command
```
signin book <desk-number> <date>
signin b <desk-number> <date>
```
 - desk-number : number of the desired desk
 - date : date when the reservation should be done

samples
```
signin book 59 20230524                          // book desk number 59 for May 24th, 2023
signin b 4 20230601                              // book desk number 4 for June 1st, 2023
signin b 32 20230601 20230602 20230603           // book desk number 32 for June 1st, 2023, June 2nd, 2023 and June 3rd, 2023
signin b 59 3 20230601                           // book desk number 9¡59 for 3 consecutive days starting June 1st, 2023
```

### cancel booking
Cancels a previouslny registered booking for a specific date

command
```
signin cancel <date>
signin c <date>
```
 - booking-id : ID of the booking to cancel

samples
```
signin cancel 20230601       // cancel booking for June 1st, 2023
signin c 20230601            // cancel booking for June 1st, 2023
```
[back to top](#signin)

### list free spaces for a given date
list all free desks for a given date

command
```
signin list-free <date>
signin lf <date>
```

 - date : date to get the free spaces

samples
```
signin list-free 20230524          // list free spaces for May 24th, 2023
signin lf 20230601                 // list free spaces for June 1st, 2023
```
[back to top](#signin)


### list bookings up to a given date
list all fbookings from the current date up to a given date

command
```
signin list-bookings <date>
signin lb <date>
```
 - date : limit date to get the list of bookings

samples
```
signin list-bookings 20230524          // list bookings from now until  May 24th, 2023
signin lb 20230601                     // list bookings from now until June 1st, 2023
```
[back to top](#signin)


### config
Display the current configuration

command
```
signin config
```
[back to top](#signin)

### config bearer
Set the signin app bearer in the config

command
```
signin config bearer <bearer>
```
 - bearer : sigin app bearer

samples
```
signin config bearer my-app-bearer          // set the bearer as my-app-bearer
[back to top](#signin)
```
