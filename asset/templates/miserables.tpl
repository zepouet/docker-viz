{{define "content"}}
<script>
    d3.json("/json/miserables", function(miserables) {
        var matrix = [],
                nodes = miserables.nodes,
                n = nodes.length;

        var margin = {top: 150, right: 20, bottom: 10, left: 150},
                width = 12 * n,
                height = 12 * n;

        var x = d3.scale.ordinal().rangeBands([0, width]),
                z = d3.scale.linear().domain([0, 4]).clamp(true),
                c = d3.scale.category10().domain(d3.range(10));

        var svg = d3.select(".main").append("svg")
                .attr("width", width + margin.left + margin.right)
                .attr("height", height + margin.top + margin.bottom)
                .append("g")
                .attr("transform", "translate(" + margin.left + "," + margin.top + ")");



        // Compute index per node.
        nodes.forEach(function(node, i) {
            node.index = i;
            node.count = 0;
            matrix[i] = d3.range(n).map(function(j) { return {x: j, y: i, z: 0}; });
        });

        // Convert links to matrix; count character occurrences.
        miserables.links.forEach(function(link) {
            matrix[link.source][link.target].z += link.value;
            //matrix[link.target][link.source].z += link.value;
        });

        // Precompute the orders.
        var orders = {
            name: d3.range(n).sort(function(a, b) { return d3.ascending(nodes[a].name, nodes[b].name); })
        };

        // The default sort order.
        x.domain(orders.name);

        svg.append("rect")
                .attr("class", "background")
                .attr("width", width)
                .attr("height", height);

        var row = svg.selectAll(".row")
                .data(matrix)
                .enter().append("g")
                .attr("class", "row")
                .attr("transform", function(d, i) { return "translate(0," + x(i) + ")"; })
                .each(row);

        row.append("line")
                .attr("x2", width);

        row.append("text")
                .attr("x", -6)
                .attr("y", x.rangeBand() / 2)
                .attr("dy", ".32em")
                .attr("text-anchor", "end")
                .attr("class", "text")
                .text(function(d, i) { return nodes[i].name; });

        var column = svg.selectAll(".column")
                .data(matrix)
                .enter().append("g")
                .attr("class", "column")
                .attr("transform", function(d, i) { return "translate(" + x(i) + ")rotate(-90)"; });

        column.append("line")
                .attr("x1", -width);

        column.append("text")
                .attr("x", 6)
                .attr("y", x.rangeBand() / 2)
                .attr("dy", ".32em")
                .attr("text-anchor", "start")
                .attr("class", "text")
                .text(function(d, i) { return nodes[i].name; });

        function row(row) {
            var cell = d3.select(this).selectAll(".cell")
                    .data(row.filter(function(d) { return d.z; }))
                    .enter().append("rect")
                    .attr("class", "cell")
                    .attr("x", function(d) { return x(d.x); })
                    .attr("width", x.rangeBand())
                    .attr("height", x.rangeBand())
                    .style("fill", function(d) { return "rgb("+d.z+",100,100)" })
                    .on("mouseover", mouseover)
                    .on("mouseout", mouseout);
        }

        function mouseover(p) {
            d3.selectAll(".row text").classed("active", function(d, i) { return i == p.y; });
            d3.selectAll(".column text").classed("active", function(d, i) { return i == p.x; });
        }

        function mouseout() {
            d3.selectAll("text").classed("active", false);
        }
    });

</script>
{{end}}