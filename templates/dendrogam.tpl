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
        var width = 1000,
            height = 1000;

        var cluster = d3.layout.cluster()
            .size([height, width - 200]);

        var diagonal = d3.svg.diagonal()
            .projection(function (d) {
                return [d.y, d.x];
            });

        var svg = d3.select("body").append("svg")
            .attr("width", width)
            .attr("height", height)
            .append("g")
            .attr("transform", "translate(50,0)");

        d3.json("/flare/{{ .type }}/json", function (error, root) {
            var nodes = cluster.nodes(root),
                links = cluster.links(nodes);

            var link = svg.selectAll(".link")
                .data(links)
                .enter().append("path")
                .attr("class", "link")
                .attr("d", diagonal);

            var node = svg.selectAll(".node")
                .data(nodes)
                .enter().append("g")
                .attr("class", "node")
                .attr("transform", function (d) {
                    return "translate(" + d.y + "," + d.x + ")";
                })

            node.append("circle")
                .attr("r", 4);

            node.append("text")
                .attr("dx", function (d) {
                    return d.children ? -6 : 8;
                })
                .attr("dy", 3)
                .style("text-anchor", function (d) {
                    return d.children ? "end" : "start";
                })
                .text(function (d) {
                    return d.name;
                });
        });
        d3.select(self.frameElement).style("height", height + "px");
    }
</script>
{{end}}