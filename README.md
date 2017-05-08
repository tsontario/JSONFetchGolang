<h1>Backend Technical Challenge</h1>

<p>This is a small program to gather, process, and sort specific JSON data relating to fulfilling orders for a business.
The code performs the following operations:</p>

<ul>
	<li>Retrieve all current orders from the web</li>
	<li>Fulfill all orders that don't contain cookies</li>
	<li>For all unfulfilled orders, place in a max heap ordered by number of cookies. Order by lowest customer id if number of cookies is the same.</li>
	<li>Iterating through the heap in an attempt to fulfill orders</li>
	<li>Once done, generate JSON output detailing the number of cookies remaining and the ids of all unfulfilled orders (sorted by customer id)</li>
</ul>	

<h2>Approach</h2>
<p>orderstructs.go contains type definitions for the relevant JSON properties we need from the website. While Go allows for anonymous parsing through JSON, specifying structs allows for greater flexibility and extensibility should requirements change in the future</p>
<p>priorityqueue.go contains an implementation of the heap interface and is used to organize orders by the above mentioned criteria. The sorting mechanism can be changed simply by changing the Less function. There currently exists a bug whereby heap.init() does not properly heapify an unordered array of orders. Consequently, the heap is currently generated in the less-efficient manner of linearly adding orders to the heap and upheaping with every addition.</p>
<p>prioritizeorders.go contains the main logic and consists of connecting to the web, building up the JSON objects and then iterating over and obtaining the output. The code is organized in such a way that it is easy to change the parameters of both the source URL and the product to be ordered by.</p>