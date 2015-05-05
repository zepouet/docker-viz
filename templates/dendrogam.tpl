{{define "content"}}
<script>
    $(document).ready(function() {
        loadD3JSon()
        setInterval(reloadD3JSon, 60000);
    });

    function reloadD3JSon() {
            $('svg').first().remove()

            loadD3JSon()
    }

    function loadD3JSon() {
        // Get JSON data
        treeJSON = d3.json("/json/{{ .type }}/dendrogam", function(error, treeData) {

            // Calculate total nodes, max label length
            var totalNodes = 0;
            var maxLabelLength = 0;
            // Misc. variables
            var i = 0;
            var duration = 750;
            var root;

            // size of the diagram
            var viewerWidth = $(document).width() - 50;
            var viewerHeight = $(document).height() - 50;

            var tree = d3.layout.tree()
                    .size([viewerHeight, viewerWidth]);

            // define a d3 diagonal projection for use by the node paths later on.
            var diagonal = d3.svg.diagonal()
                    .projection(function(d) {
                        return [d.y, d.x];
                    });

            // A recursive helper function for performing some setup by walking through all nodes

            function visit(parent, visitFn, childrenFn) {
                if (!parent) return;

                visitFn(parent);

                var children = childrenFn(parent);
                if (children) {
                    var count = children.length;
                    for (var i = 0; i < count; i++) {
                        visit(children[i], visitFn, childrenFn);
                    }
                }
            }

            // Call visit function to establish maxLabelLength
            visit(treeData, function(d) {
                totalNodes++;
                maxLabelLength = Math.max(d.name.length, maxLabelLength);

            }, function(d) {
                return d.children && d.children.length > 0 ? d.children : null;
            });


            // sort the tree according to the node names

            function sortTree() {
                tree.sort(function(a, b) {
                    return b.name.toLowerCase() < a.name.toLowerCase() ? 1 : -1;
                });
            }
            // Sort the tree initially incase the JSON isn't in a sorted order.
            sortTree();


            function zoom() {
                svgGroup.attr("transform", "translate(" + d3.event.translate + ")scale(" + d3.event.scale + ")");
            }

            // define the zoomListener which calls the zoom function on the "zoom" event constrained within the scaleExtents
            var zoomListener = d3.behavior.zoom().scaleExtent([0.1, 3]).on("zoom", zoom);

            // define the baseSvg, attaching a class for styling and the zoomListener
            var baseSvg = d3.select("#tree-container").append("svg")
                    .attr("width", viewerWidth)
                    .attr("height", viewerHeight)
                    .attr("class", "overlay")
                    .call(zoomListener);

            dragListener = d3.behavior.drag();

            function centerNode(source) {
                scale = zoomListener.scale();
                x = -source.y0;
                y = -source.x0;
                x = scale + 100;
                y = scale;
                d3.select('g').transition()
                        .duration(duration)
                        .attr("transform", "translate(" + x + "," + y + ")scale(" + scale + ")");
                zoomListener.scale(scale);
                zoomListener.translate([x, y]);
            }

            function click(d) {
                if (d3.event.defaultPrevented) return; // click suppressed
                d = toggleChildren(d);
                update(d);
                centerNode(d);
            }

            function update(source) {
                var levelWidth = [1];
                var childCount = function(level, n) {

                    if (n.children && n.children.length > 0) {
                        if (levelWidth.length <= level + 1) levelWidth.push(0);

                        levelWidth[level + 1] += n.children.length;
                        n.children.forEach(function(d) {
                            childCount(level + 1, d);
                        });
                    }
                };

                childCount(0, root);
                var newHeight = d3.max(levelWidth) * 25; // 25 pixels per line
                tree = tree.size([newHeight, viewerWidth]);

                // Compute the new tree layout.
                var nodes = tree.nodes(root).reverse(),
                        links = tree.links(nodes);

                // Set widths between levels based on maxLabelLength.
                nodes.forEach(function(d) {
                    d.y = (d.depth * (maxLabelLength * 10));
                });

                // Update the nodes…
                node = svgGroup.selectAll("g.node")
                        .data(nodes, function(d) {
                            return d.id || (d.id = ++i);
                        });

                // Enter any new nodes at the parent's previous position.
                var nodeEnter = node.enter().append("g")
                        .call(dragListener)
                        .attr("class", "node")
                        .attr("transform", function(d) {
                            return "translate(" + source.y0 + "," + source.x0 + ")";
                        })
                        .on('click', click);

                nodeEnter.append("circle")
                        .attr('class', 'nodeCircle')
                        .attr("r", 0)
                        .style("fill", function(d) {
                            return d._children ? "lightsteelblue" : "#fff";
                        });

                nodeEnter.append("text")
                        .attr("x", function(d) {
                            return d.children || d._children ? -10 : 10;
                        })
                        .attr("dy", ".35em")
                        .attr('class', 'nodeText')
                        .attr("text-anchor", function(d) {
                            return d.children || d._children ? "end" : "start";
                        })
                        .text(function(d) {
                            return d.name;
                        })
                        .style("fill-opacity", 0);

                // Update the text to reflect whether node has children or not.
                node.select('text')
                        .attr("x", function(d) {
                            return d.children || d._children ? -10 : 10;
                        })
                        .attr("text-anchor", function(d) {
                            return d.children || d._children ? "end" : "start";
                        })
                        .text(function(d) {
                            return d.name;
                        });

                // Change the circle fill depending on whether it has children and is collapsed
                node.select("circle.nodeCircle")
                        .attr("r", 4.5)
                        .style("fill", function(d) {
                            return d._children ? "lightsteelblue" : "#fff";
                        });

                // Transition nodes to their new position.
                var nodeUpdate = node.transition()
                        .duration(duration)
                        .attr("transform", function(d) {
                            return "translate(" + d.y + "," + d.x + ")";
                        });

                // Fade the text in
                nodeUpdate.select("text")
                        .style("fill-opacity", 1);

                // Update the links…
                var link = svgGroup.selectAll("path.link")
                        .data(links, function(d) {
                            return d.target.id;
                        });

                // Enter any new links at the parent's previous position.
                link.enter().insert("path", "g")
                        .attr("class", "link")
                        .attr("d", function(d) {
                            var o = {
                                x: source.x0,
                                y: source.y0
                            };
                            return diagonal({
                                source: o,
                                target: o
                            });
                        });

                // Transition links to their new position.
                link.transition()
                        .duration(duration)
                        .attr("d", diagonal);
            }

            // Append a group which holds all nodes and which the zoom Listener can act upon.
            var svgGroup = baseSvg.append("g");

            // Define the root
            root = treeData;
            root.x0 = viewerHeight / 2;
            root.y0 = 0;

            // Layout the tree initially and center on the root node.
            update(root);
            centerNode(root);
        });
    }
</script>
<div id="tree-container"></div>
{{end}}