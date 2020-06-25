9.3: Perform Regular Automated Port Scans
=========================================================
Perform automated port scans on a regular basis against all systems and alert if unauthorized ports are detected on a system.

.. list-table::
	:header-rows: 1

	* - Asset Type
	  - Security Function
	  - Implementation Groups
	* - Devices
	  - Detect
	  - 2, 3

Dependencies
------------
* Sub-control 2.5: Integrate Software and Hardware Asset Inventories

Inputs
------
#. t(i): the timestamp at which a port scan i has been performed
#. N: the number of port scans (timestamps) taken so far
#. M: the maximum possible irregularity (can be fixed as 30 day)
#. T: (optional) target/desirable review interval threshold
#. D: the number of port scan in which at least one anomaly was detected
#. L: The total number of port scans
#. UP: The number of alerts received due to unauthorized ports (M5)
#. NP: The number of unauthorized ports (M6)

Operations
----------
* Enumerate endpoints and identify port scanning software
* Calculate measures M1 - M6 for each port scanning software, tracking endpoints covered
* Enumerate set of endpoints covered by port scanning software
* Compare enumeration of covered endpoints against the list of all endpoints to identify those endpoints that are not covered

Measures
--------
* M1 (the average of port scans) = :code:`SUM from i=1..N  ( t(i+1) - t(i) ) / N`
* M2 (Regularity Measure of Port Scan) = :code:`(SUM from i=1..N  ( (t(i+1) - t(i) - M1)^2 / N ) / M`
* M3 (Threshold-based Regularity Measure of Port Scan) = :code:`(SUM from i=1..N ( ( t(i+1) - t(i) - T )^2 / N ) / M`
* M4 (The Probability of detecting an anomaly in port scans) = :code:`D / L`
* M5 = Count of alerts received due to unauthorized ports
* M6 = Count of unauthorized ports
* M7 = List of endpoints covered by port scanning tools
* M8 = List of endpoints not covered by port scanning tools
* M9 = Count of endpoints covered by port scanning tools
* M10 = Count of endpoints

Metrics
-------

Quality of Port Scan
^^^^^^^^^^^^^^^^^^^^
.. list-table::

	* - **Metric**
	  - | Quality of review is high if and only if the review is highly regular and the potential for detecting anomalies (at least one per review) is also high.
	* - **Calculation**
	  - :code:`(1 - M2) * M4` or (if M3) :code:`(1 - M3) * M4`

Quality
^^^^^^^
.. list-table::

	* - **Metric**
	  - | Ratio of unauthorized ports reported
	* - **Calculation**
	  - :code:`M5 / M6`

Coverage
^^^^^^^^
.. list-table::

	* - **Metric**
	  - | Ratio of covered endpoints to the total number of endpoints
	* - **Calculation**
	  - :code:`M9 / M10`

.. history
.. authors
.. license
