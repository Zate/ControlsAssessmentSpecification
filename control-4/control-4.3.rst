4.3: Ensure the Use of Dedicated Administrative Accounts
=========================================================
Ensure that all users with administrative account access use a dedicated or secondary account for elevated activities. This account should only be used for administrative activities and not Internet browsing, email, or similar activities.

.. list-table::
	:header-rows: 1

	* - Asset Type
	  - Security Function
	  - Implementation Groups
	* - Users
	  - Protect
	  - 1, 2, 3

Dependencies
------------
* None

Inputs
------
#. The list of users defined as Administrators
#. The list of user accounts for the users defined in Input 1
#. The list of users NOT defined as Administrators
#. The list of user accounts for the users defined in Input 3
#. The list of all user accounts.
#. The list of all Administrative user accounts
#. The list of non-Administrative user accounts

Operations
----------
#. For each user defined in Input 1, collect the Administrative user account for that user from Input 6 and the non-Administrative user account from Input 7
#. For each user defined in Input 3, collect any Administrative user account for that user from Input 6 and the non-Administrative user account from Input 7

Measures
--------
* M1 = The list of defined Administrative users
* M2 = The count of M1
* M3 = The list of users collected in Operation 1
* M4 = The count of M3
* M5 = The list of users collected in Operation 2
* M6 = The count of M5
* M7 = The number of users defined as Administrators
* M8 = For each user, this measure describes the number of user accounts identified by Operation 1
* M9 = For each user, this measure describes the number of user accounts identified by Operation 2


Metrics
-------

Administrative User Accounts
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
.. list-table::

	* - **Metric**
	  - | This metric is intended to determine whether those users identified as Administrative-level have at least one Administrative-level and one non-Administrative level user account. The mapping performed by Operation 1 must show that, for each Administrative-level user, at least 1 Administrative-level user account and at least 1 non-Administrative-level user account are available.
	* - **Calculation**
	  - :code:`M8 = M1`

Unauthorized User Accounts
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
.. list-table::

	* - **Metric**
	  - | This metric is intended to illustrate any non-Administrative-level users that have been assigned an Administrative-level user account.
	* - **Calculation**
	  - If :code:`M6 < 1`

.. history
.. authors
.. license
