19.6: Publish Information Regarding Reporting Computer Anomalies and Incidents
==============================================================================
Publish information for all workforce members, regarding reporting computer anomalies and incidents, to the incident handling team.  Such information should be included in routine employee awareness activities.

.. list-table::
	:header-rows: 1

	* - Asset Type
	  - Security Function
	  - Implementation Groups
	* - N/A
	  - N/A
	  - 1, 2, 3

Dependencies
------------
* Subcontrol 19.1: Document Incident Response Procedures

Inputs
-----------
#. Incident response plan
#. Security awareness program documentation

Operations
----------
#. Determine whether incident response plan exists (becomes M1)
#. Determine whether the security awareness documentation exists (becomes M2)
#. If both exist, then review the security awareness plan (determine M3 and M4)

Measures
--------
* M1 = Boolean value indicating whether an incident response plan exists; 1 if an incident response plan exists, 0 otherwise.
* M2 = Boolean value indicating whether a security awareness program exists; 1 if an incident response plan exists, 0 otherwise.
* M3 = The incident response plan requires publishing incident reporting information for all workforce members as part of the organization's security awareness program
* M4 = The security awareness program publishes incident reporting information for all workforce members

Metrics
-------

Coverage
^^^^^^^^
.. list-table::

	* - **Metric**
	  - | Is information regarding reporting of computer anomalies and incidents published for all workforce members?
	* - **Calculation**
	  - :code:`M1 AND M2 AND M3 AND M4`

.. history
.. authors
.. license
