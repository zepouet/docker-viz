{{define "content"}}
<script>
    $(document).ready(function() {
        loadD3JSon()
        setInterval(reloadD3JSon, 10000);
    });

    function reloadD3JSon() {
        $('body').fadeOut("slow", function() {
            loadD3JSon()
            $('svg').first().remove()
            $('body').fadeIn("slow")
        })
    }

    function loadD3JSon() {
        var diameter = 1000,
            format = d3.format(",d"),
            color = d3.scale.category20c();

        var bubble = d3.layout.pack()
            .sort(null)
            .size([diameter, diameter])
            .padding(1.5);

        var svg = d3.select("body").append("svg")
            .attr("width", diameter)
            .attr("height", diameter)
            .attr("class", "bubble");

        d3.json("/flare/{{ .type }}/json", function(error, root) {
            var node = svg.selectAll(".node")
                .data(bubble.nodes(classes(root))
                        .filter(function(d) { return !d.children; }))
                .enter().append("g")
                .attr("class", "node")
                .attr("transform", function(d) { return "translate(" + d.x + "," + d.y + ")"; });

            node.append("title")
                .text(function(d) { return d.className + ": " + format(d.value); });

            node.append("circle")
                .attr("r", function(d) { return d.r; })
                .style("fill", function(d) { return color(d.packageName); });

            node.append("text")
                .attr("dy", ".3em")
                .style("text-anchor", "middle")
                .text(function(d) { return d.className.substring(0, d.r / 3); });
        });

        // Returns a flattened hierarchy containing all leaf nodes under the root.
        function classes(root) {
            var classes = [];

            function recurse(name, node) {
                if (node.children) node.children.forEach(function(child) { recurse(node.name, child); });
                else classes.push({packageName: name, className: node.name, value: node.size});
            }

            recurse(null, root);
            return {children: classes};
        }

        d3.select(self.frameElement).style("height", diameter + "px");
    }
</script>
{{end}}