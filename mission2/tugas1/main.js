const availableItems = document.querySelectorAll('.product');
const btnProd = document.querySelectorAll('.btn-prod');
const cartCont = document.querySelector('.cart-cont');
const total = document.querySelector('.total');
const cartItems = [];

// Fetch data and display products
fetch("data.json")
    .then(response => response.json())
    .then(data => {
        displayData(data);
    })
    .catch((error) => {
        console.error('Error fetching data: ', error);
    });

function formatRupiah(angka) {
    var reverse = angka.toString().split('').reverse().join('');
    var ribuan = reverse.match(/\d{1,3}/g);
    ribuan = ribuan.join('.').split('').reverse().join('');
    return 'Rp ' + ribuan;
}

function displayData(data) {
    const catProducts = document.querySelector('.cat-products');

    data.forEach((product, index) => {
        const productElement = document.createElement('div');
        productElement.className = 'product col-md-4';
        productElement.innerHTML = `
            <img class="imgProduct" src="${product.img}">
            <h3>${product.title}</h3>
            <h5>${formatRupiah(parseFloat(product.price))}</h5>
            <div class="order align-items-end">
                <button class="min col-md-2 btn btn-secondary">-</button>
                <input class="input-value col-md-3 text-center" type="text" value="0">
                <button class="plus col-md-2 btn btn-secondary">+</button>
            </div>
            <button class="btn-prod item-1 btn btn-success" data-product-id="${index}">Add</button>
        `;

        catProducts.appendChild(productElement);

        // Add event listener to the "Add" button
        const addButton = productElement.querySelector('.btn-prod');
        addButton.addEventListener('click', () => {
            const productId = addButton.getAttribute('data-product-id');
            const selectedProduct = data[productId];

            // Get the quantity from the input field
            const quantity = parseInt(productElement.querySelector('.input-value').value);

            if (quantity > 0) {
                // Check if the selected product is already in cartItems
                const existingCartItemIndex = cartItems.findIndex(item => item.title === selectedProduct.title);

                if (existingCartItemIndex !== -1) {
                    // If the product is already in cartItems, update the quantity
                    cartItems[existingCartItemIndex].quantity += quantity;
                } else {
                    // If the product is not in cartItems, add it with the quantity
                    selectedProduct.quantity = quantity;
                    cartItems.push(selectedProduct);
                }

                // Display the updated cart
                displayCart();
            } else {
                alert("Barang tidak boleh kurang dari 1");
            }
        });

        // Add event listeners to the "plus" and "min" buttons
        const plusButton = productElement.querySelector('.plus');
        const minButton = productElement.querySelector('.min');
        const inputValue = productElement.querySelector('.input-value');

        plusButton.addEventListener('click', () => {
            let currentValue = parseInt(inputValue.value);
            inputValue.value = currentValue + 1;
        });

        minButton.addEventListener('click', () => {
            let currentValue = parseInt(inputValue.value);
            if (currentValue > 0) {
                inputValue.value = currentValue - 1;
            } else {
                alert("Barang tidak boleh kurang dari 1");
            }
        });
    });
}




function displayCart() {
    // Clear the cart content before adding new items
    cartCont.innerHTML = '';

    let totalPriceBeforeTax = 0;
    let totalTax = 0;

    cartItems.forEach((item) => {
        const price = Number(item.price);
        const quantity = item.quantity

        totalPriceBeforeTax += price * quantity;

        cartCont.innerHTML += `
            <div class="cartItems row">
                <div class="col-md-3">
                    <img class="imgMyCart" src=${item.img}>
                </div>
                <div class="col-md-4 text-start">
                    <span>${item.title}</span>
                </div>
                <div class="col-md-3 text-start">
                    <span>${formatRupiah(price)}</span>
                </div>
                <div class="col-md-2">
                    <span>x${quantity}</span>
                </div>
            </div>
        `;
    });

    // Calculate the tax and total price after tax
    totalTax = totalPriceBeforeTax * 0.11;
    const totalPriceAfterTax = totalPriceBeforeTax + totalTax;

    // Update the total with the new values and format them as Indonesian currency
    total.innerHTML = `
        <div class="col-md-7"></div>
        <div class="totalPrice col-md-5">
            <div>Total Pembelian: ${formatRupiah(totalPriceBeforeTax)}</div>
            <div>Pajak 11%: ${formatRupiah(totalTax)}</div>
            <div>Total Bayar: ${formatRupiah(totalPriceAfterTax)}</div> 
            <button class="btn btn-primary btn-print-receipt mt-1 mb-1">Cetak Struk Pembayaran</button>
        </div>
        
    `;
}

