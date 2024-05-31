package resolver

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/spf13/afero"
)

// Helper function to initialize a DependencyResolver with sample data.
func setupTestResolver() *DependencyResolver {
	fs := afero.NewMemMapFs() // Using an in-memory filesystem for testing
	logger := log.New(nil)
	resolver, err := NewDependencyResolver(fs, logger)
	if err != nil {
		log.Fatalf("Failed to create dependency resolver: %v", err)
	}

	resolver.Resources = []ResourceEntry{
		{Resource: "a", Name: "A", Sdesc: "Resource A", Ldesc: "The first resource in the alphabetical order", Category: "example", Requires: []string{}},
		{Resource: "b", Name: "B", Sdesc: "Resource B", Ldesc: "The second resource, dependent on A", Category: "example", Requires: []string{"a"}},
		{Resource: "c", Name: "C", Sdesc: "Resource C", Ldesc: "The third resource, dependent on B", Category: "example", Requires: []string{"b"}},
		{Resource: "d", Name: "D", Sdesc: "Resource D", Ldesc: "The fourth resource, dependent on C", Category: "example", Requires: []string{"c"}},
		{Resource: "e", Name: "E", Sdesc: "Resource E", Ldesc: "The fifth resource, dependent on D", Category: "example", Requires: []string{"d"}},
		{Resource: "f", Name: "F", Sdesc: "Resource F", Ldesc: "The sixth resource, dependent on E", Category: "example", Requires: []string{"e"}},
		{Resource: "g", Name: "G", Sdesc: "Resource G", Ldesc: "The seventh resource, dependent on F", Category: "example", Requires: []string{"f"}},
		{Resource: "h", Name: "H", Sdesc: "Resource H", Ldesc: "The eighth resource, dependent on G", Category: "example", Requires: []string{"g"}},
		{Resource: "i", Name: "I", Sdesc: "Resource I", Ldesc: "The ninth resource, dependent on H", Category: "example", Requires: []string{"h"}},
		{Resource: "j", Name: "J", Sdesc: "Resource J", Ldesc: "The tenth resource, dependent on I", Category: "example", Requires: []string{"i"}},
		{Resource: "k", Name: "K", Sdesc: "Resource K", Ldesc: "The eleventh resource, dependent on J", Category: "example", Requires: []string{"j"}},
		{Resource: "l", Name: "L", Sdesc: "Resource L", Ldesc: "The twelfth resource, dependent on K", Category: "example", Requires: []string{"k"}},
		{Resource: "m", Name: "M", Sdesc: "Resource M", Ldesc: "The thirteenth resource, dependent on L", Category: "example", Requires: []string{"l"}},
		{Resource: "n", Name: "N", Sdesc: "Resource N", Ldesc: "The fourteenth resource, dependent on M", Category: "example", Requires: []string{"m"}},
		{Resource: "o", Name: "O", Sdesc: "Resource O", Ldesc: "The fifteenth resource, dependent on N", Category: "example", Requires: []string{"n"}},
		{Resource: "p", Name: "P", Sdesc: "Resource P", Ldesc: "The sixteenth resource, dependent on O", Category: "example", Requires: []string{"o"}},
		{Resource: "q", Name: "Q", Sdesc: "Resource Q", Ldesc: "The seventeenth resource, dependent on P", Category: "example", Requires: []string{"p"}},
		{Resource: "r", Name: "R", Sdesc: "Resource R", Ldesc: "The eighteenth resource, dependent on Q", Category: "example", Requires: []string{"q"}},
		{Resource: "s", Name: "S", Sdesc: "Resource S", Ldesc: "The nineteenth resource, dependent on R", Category: "example", Requires: []string{"r"}},
		{Resource: "t", Name: "T", Sdesc: "Resource T", Ldesc: "The twentieth resource, dependent on S", Category: "example", Requires: []string{"s"}},
		{Resource: "u", Name: "U", Sdesc: "Resource U", Ldesc: "The twenty-first resource, dependent on T", Category: "example", Requires: []string{"t"}},
		{Resource: "v", Name: "V", Sdesc: "Resource V", Ldesc: "The twenty-second resource, dependent on U", Category: "example", Requires: []string{"u"}},
		{Resource: "w", Name: "W", Sdesc: "Resource W", Ldesc: "The twenty-third resource, dependent on V", Category: "example", Requires: []string{"v"}},
		{Resource: "x", Name: "X", Sdesc: "Resource X", Ldesc: "The twenty-fourth resource, dependent on W", Category: "example", Requires: []string{"w"}},
		{Resource: "y", Name: "Y", Sdesc: "Resource Y", Ldesc: "The twenty-fifth resource, dependent on X", Category: "example", Requires: []string{"x"}},
		{Resource: "z", Name: "Z", Sdesc: "Resource Z", Ldesc: "The twenty-sixth resource, dependent on Y", Category: "example", Requires: []string{"y"}},
	}
	for _, entry := range resolver.Resources {
		resolver.resourceDependencies[entry.Resource] = entry.Requires
	}
	return resolver
}

func TestShowResourceEntry(t *testing.T) {
	resolver := setupTestResolver()

	// Capture the output
	var output strings.Builder
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	resolver.ShowResourceEntry("a")

	w.Close()
	os.Stdout = old
	io.Copy(&output, r)

	expectedOutput := "Resource: a\nName: A\nShort Description: Resource A\nLong Description: The first resource in the alphabetical order\nCategory: example\nRequirements: []\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}

func TestListDirectDependencies(t *testing.T) {
	resolver := setupTestResolver()

	// Capture the output
	var output strings.Builder
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	resolver.Graph.ListDirectDependencies("z")

	w.Close()
	os.Stdout = old
	io.Copy(&output, r)

	expectedOutput := `z
z -> y
z -> y -> x
z -> y -> x -> w
z -> y -> x -> w -> v
z -> y -> x -> w -> v -> u
z -> y -> x -> w -> v -> u -> t
z -> y -> x -> w -> v -> u -> t -> s
z -> y -> x -> w -> v -> u -> t -> s -> r
z -> y -> x -> w -> v -> u -> t -> s -> r -> q
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h -> g
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h -> g -> f
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h -> g -> f -> e
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h -> g -> f -> e -> d
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h -> g -> f -> e -> d -> c
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h -> g -> f -> e -> d -> c -> b
z -> y -> x -> w -> v -> u -> t -> s -> r -> q -> p -> o -> n -> m -> l -> k -> j -> i -> h -> g -> f -> e -> d -> c -> b -> a
`
	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}

func TestListDependencyTree(t *testing.T) {
	resolver := setupTestResolver()

	// Capture the output
	var output strings.Builder
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	resolver.Graph.ListDependencyTree("z")

	w.Close()
	os.Stdout = old
	io.Copy(&output, r)

	expectedOutput := "z <- y <- x <- w <- v <- u <- t <- s <- r <- q <- p <- o <- n <- m <- l <- k <- j <- i <- h <- g <- f <- e <- d <- c <- b <- a\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}

func TestListDependencyTreeList(t *testing.T) {
	resolver := setupTestResolver()

	// Capture the output
	var output strings.Builder
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	resolver.Graph.ListDependencyTreeTopDown("z")

	w.Close()
	os.Stdout = old
	io.Copy(&output, r)

	expectedOutput :=
		`a
b
c
d
e
f
g
h
i
j
k
l
m
n
o
p
q
r
s
t
u
v
w
x
y
z
`

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}
