<script lang="ts">
    import {onMount} from "svelte";
    import {page} from "$app/stores";

    const updatedAt = 'May 5, 2023',
        mobilAppName = 'RFX Services',
        country = 'United States',
        lawState = 'Texas',
        contactEmail = 'contact@rfx.services';

    let domain, termLink;
    onMount(() => {
        domain = `${$page.url.protocol}//${$page.url.host}`;
        termLink = `${domain}${$page.url.pathname}`;
    })

</script>

<div class="container mx-auto p-4">
    <h1 class="text-4xl font-semibold text-center">ROOFIX General Terms and Conditions for Services</h1>
    <h4 class="mb-3 mt-8 font-bold">Last Updated: {updatedAt}</h4>

    <p class="mb-3">These General Terms and Conditions for Services (“General Terms”) are applicable to any Proposal
        issued by Roofix Technologies ("Roofix") and accepted by Company for the provision of Services. Each Proposal
        issued by Roofix shall be subject to these General Terms. These General Terms, along with any onboarding
        documents executed by the parties, collectively constitute the “Agreement” by and between Roofix and Company.
        Capitalized terms have the meanings ascribed to them throughout the Agreement. These General Terms prevail
        over any of Company's general terms and conditions regardless of whether or when Company has submitted its
        request for proposal, order, or such terms. Provision of Services to Company does not constitute acceptance
        of any of Company's terms and conditions and does not serve to modify or amend these General Terms.</p>
    <p class="mb-3">IF THE INDIVIDUAL ACCEPTING THIS AGREEMENT IS ACCEPTING ON BEHALF OF A COMPANY OR OTHER LEGAL ENTITY,
        SUCH INDIVIDUAL REPRESENTS THAT THEY HAVE THE AUTHORITY TO BIND SUCH ENTITY AND ITS AFFILIATES TO THESE TERMS AND
        CONDITIONS, IN WHICH CASE THE TERM “COMPANY” SHALL REFER TO SUCH ENTITY AND ITS AFFILIATES. IF THE INDIVIDUAL
        ACCEPTING THIS AGREEMENT DOES NOT HAVE SUCH AUTHORITY, OR DOES NOT AGREE WITH THESE TERMS AND CONDITIONS, SUCH
        INDIVIDUAL MUST NOT ACCEPT THIS AGREEMENT AND MAY NOT USE THE SERVICES.</p>

    <h4 class="mb-3 mt-8 font-bold">1. DEFINITIONS</h4>
    <p class="mb-3">“Affiliate” means an entity or person that directly or indirectly, through one or more intermediaries,
        controls, is controlled by or in common control with a party to this Agreement. For purposes of this definition,
        control means direct or indirect ownership or control of 50% of the voting interest of the subject entity.
    </p>
    <p class="mb-3">“Certified Partners” means persons or entities Company has entered into separate agreements for
        provision of certain services. </p>
    <p class="mb-3">“Company” means the person or entity procuring the Services identified and described in the applicable Proposal.</p>
    <p class="mb-3">“Fees” means fees payable for the Services as set forth in a Proposal including the contract price,
        or otherwise agreed to by the parties.</p>
    <p class="mb-3">“Platform” means Roofix’ proprietary software application system and related database used in connection
        with the Services. </p>
    <p class="mb-3">“Property Owner” is the owner of the property in regard to which Company is procuring the Services.</p>
    <p class="mb-3">“Proposal” means the Roofix proposal, estimate or ordering document by which Company agrees to
        purchase the Services, containing applicable details, including the Fees, the Services selected, applicable
        timelines and other relevant information, By approving or accepting a Proposal hereunder, an Affiliate agrees
        to be bound by the terms of this Agreement as if it were an original party hereto.
    </p>

    <h4 class="mb-3 mt-8 font-bold">2. SERVICES</h4>
    <p class="mb-3">(a) Roofix shall provide the Services to Company as described in the Proposal in accordance with
        these General Terms (the “Services”). </p>
    <p class="mb-3">(b) Assumptions and Disclosures:</p>
    <ul class="list-disc ml-6 mb-3 list-outside">
        <li>Company has entered or will enter into an agreement with the Property Owner with regard to the Services.</li>
        <li>If Roofix’ performance of its obligations under this Agreement is prevented or delayed by any act or
            omission of Company (or its agents, subcontractors, consultants, or employees) or of Property Owner,
            Roofix shall not be deemed in breach of its obligations under this Agreement or otherwise liable for
            any costs, charges, or losses sustained or incurred by Customer, in each case, to the extent arising
            directly or indirectly from such prevention or delay.</li>
        <li>Incorrect materials selection and scope requirements may result in cancellation of the applicable project
            and termination of the applicable Proposal, in which case Company shall reimburse all of Roofix’ costs
            incurred by such cancellation.</li>
        <li>Low slope roofing areas will be estimated as Remove Modified Bitumen/Rolled Comp and Install New Modified
            Bitumen/Rolled Comp. Any changes to the low slope scope of work will require customized pricing and may
            result in a Change Order or revised project financials.</li>
        <li>Timelines for completion of the Services are subject to and may be delayed by matters, requirements or
            events outside of Roofix’ control including permitting requirements and weather.</li>
        <li>Roofix is unable to complete repairs and replacements on mobile homes in California.</li>
    </ul>

    <h4 class="mb-3 mt-8 font-bold">3. PERFORMANCE DATES</h4>
    <p class="mb-3">Roofix shall use reasonable efforts to meet any performance dates specified in the Proposal, and
        any such dates shall be estimates only.</p>

    <h4 class="mb-3 mt-8 font-bold">4. COMPANY'S OBLIGATIONS</h4>
    <p class="mb-3">Company shall:</p>
    <p class="mb-3">(a) cooperate with Roofix in all matters relating to the Services and ensuring Roofix rights to
        access to Property Owner’s premises, as may reasonably be requested by Roofix, for the purposes of performing
        the Services;</p>
    <p class="mb-3">(b) respond promptly to any Roofix request to provide direction, information, approvals,
        authorizations, or decisions that are reasonably necessary for Roofix to perform Services in accordance with
        the requirements of this Agreement;</p>
    <p class="mb-3">(c) provide such materials or information as Roofix may request to carry out the Services in a
        timely manner and ensure that such Company materials or information are complete and accurate in all material
        respects including without limitation the estimated square footage of the project; and</p>
    <p class="mb-3">(d) obtain and maintain all necessary licenses and consents and comply with all applicable laws in
        relation to the Services before the date on which the Services are to start (except for roofing permits and
        licenses which are Roofix’ responsibility with regard to the Services); and</p>
    <p class="mb-3">(e) Obtain affirmative consent from Property Owner to permit Company to communicate with Property
        Owner regarding the Services including via email, telephone calls and text messages.</p>

    <h4 class="mb-3 mt-8 font-bold">5. COMPANY'S ACTS OR OMISSIONS</h4>
    <p class="mb-3">If Roofix' performance of its obligations under this Agreement is prevented or delayed by any act
        or omission of Company or its agents, subcontractors, consultants, or employees, Roofix shall not be deemed in
        breach of its obligations under this Agreement or otherwise liable for any costs, charges, or losses sustained
        or incurred by Company, in each case, to the extent arising directly or indirectly from such prevention or delay.</p>

    <h4 class="mb-3 mt-8 font-bold">6. CHANGE ORDERS</h4>
    <p class="mb-3">If either party wishes to change the scope or performance of the Services, it shall submit details
        of the requested change to the other party in writing (a “Change Order”). Neither party shall be bound by any
        Change Order unless mutually agreed upon in accordance with Section 25.</p>

    <h4 class="mb-3 mt-8 font-bold">7. FEES AND EXPENSES; PAYMENT TERMS; INTEREST ON LATE PAYMENTS</h4>
    <p class="mb-3">(a) In consideration of the provision of the Services by Roofix and the rights granted to Company
        under this Agreement, Company shall pay the Fees set forth in the Proposal. Company shall pay all invoiced
        amounts due to Roofix as set forth in the Proposal, on Roofix invoice or subject to payment terms agreed to
        by the parties. Company shall make all payments hereunder in US dollars by wire transfer, check, or other
        payment method acceptable to Roofix. FEES SET FORTH IN THE APPLICABLE PROPOSAL ARE NON-NEGOTIABLE AND ARE
        BINDING UPON ACCEPTANCE OF THE PROPOSAL BY COMPANY.</p>
  <p class="mb-3">(b) In the event payments are not received by Roofix after becoming due, Roofix may: (i) charge
    interest on any such unpaid amounts at a rate of 1.5% per month or, if lower, the maximum amount permitted under
    applicable law, from the date such payment was due until the date paid; and (ii) suspend performance for all
    Services being provided to Company and place Company’s account “on hold” (including on the Platform) until
    payment has been made in full.</p>
  <p class="mb-3">(c) COMPANY UNDERSTANDS THAT THE PROPERTY OWNER IS ULTIMATELY RESPONSIBLE FOR ALL PAYMENTS TO ROOFIX
    AND WILL BE SUBJECT TO LIENS, AS PERMITTED UNDER APPLICABLE LAW, AND LITIGATION IF PAYMENT IS NOT TENDERED ACCORDINGLY.</p>

  <h4 class="mb-3 mt-8 font-bold">8. TAXES</h4>
  <p class="mb-3">Company shall be responsible for all sales, use and excise taxes, and any other similar taxes,
    duties and charges of any kind imposed by any federal, state, or local governmental entity on any amounts payable
    by Company hereunder.</p>

  <h4 class="mb-3 mt-8 font-bold">9. INTELLECTUAL PROPERTY</h4>
  <p class="mb-3">All intellectual property rights, including copyrights, patents, patent disclosures and inventions
    (whether patentable or not), trademarks, service marks, trade secrets, know-how and other confidential information,
    trade dress, trade names, logos, corporate names, and domain names, together with all of the goodwill associated
    therewith, derivative works and all other rights (collectively, “Intellectual Property Rights”) in and to all
    documents, work product, and other materials that are delivered to Company under this Agreement or prepared by or
    on behalf of Roofix in the course of performing the Services, including any items identified as such in the
    Proposal (collectively, the “Deliverables”) except for any Confidential Information of Company or Company
    materials shall be owned by Roofix. Roofix hereby grants Company a license to use all Intellectual Property
    Rights free of additional charge and on a non-exclusive, worldwide, non-transferable, non-sublicensable,
    fully paid-up, royalty-free, and perpetual basis to the extent necessary to enable Company to make reasonable
    use of the Deliverables and the Services.</p>

  <h4 class="mb-3 mt-8 font-bold">10. CONFIDENTIAL INFORMATION</h4>
  <p class="mb-3">(a) All non-public, confidential or proprietary information of Roofix, including, but not limited to,
    trade secrets, technology, information pertaining to business operations and strategies, and information pertaining
    to Company, pricing, and marketing (collectively, “Confidential Information”), disclosed by Roofix to Company,
    whether disclosed orally or disclosed or accessed in written, electronic or other form or media, and whether or
    not marked, designated or otherwise identified as "confidential," in connection with the provision of the Services
    and this Agreement is confidential, and shall not be disclosed or copied by Company without the prior written
    consent of Roofix. Confidential Information does not include information that is: (i) in the public domain;
    (ii) known to Company at the time of disclosure; or (iii) rightfully obtained by Company on a non-confidential
    basis from a third party.</p>
  <p class="mb-3">(b) Company agrees to use the Confidential Information only to make use of the Services and
    Deliverables.</p>
  <p class="mb-3">(c) Roofix shall be entitled to injunctive relief for any violation of this Section.</p>

  <h4 class="mb-3 mt-8 font-bold">11. REPRESENTATION AND WARRANTY</h4>
  <p class="mb-3">(a) Roofix represents and warrants to Company that it shall perform the Services in a professional
    and workmanlike manner in accordance with generally recognized industry standards for similar services.</p>
  <p class="mb-3">(b) In the event Services are for roof repairs, Company will provide Property Owner a limited 90 Day
    “Workmanship Warranty” to the Property Owner.</p>
  <p class="mb-3">(c) In the event, Services are for full roof replacements, warranties shall be as follows:</p>
  <ul class="list-disc ml-6 mb-3 list-outside">
    <li>Manufacturer Warranty by shingle brand and type.</li>
    <li>Limited Lifetime Workmanship for projects fulfilled by J Wales.</li>
    <li>Minimum 5 year Workmanship for projects fulfilled by Certified Partners of Roofix.</li>
    <li>Roofix shall transfer to Property Owner all Manufacturers' written warranties upon request.</li>
  </ul>
  <p class="mb-3">(d) No warranties will be provided for the following:</p>
  <ul class="list-disc ml-6 mb-3 list-outside">
    <li>Comp Out: Remove section(s) tile and install asphalt for solar array. Back fill tile.</li>
    <li>Layover: Install new shingles or system over existing top layer of roofing material.</li>
    <li>Partial: Remove and Replace only a portion or slope(s) of the roof.</li>
    <li>Tile Repaper: Detach existing tile, install Boral Tile Seal, and reset original tile.</li>
    <li>Solar Penetrations: Areas where solar work has penetrated or impacted the roof system viability.</li>
  </ul>
  <p class="mb-3">(e) Roofix shall not be liable for a breach of the warranty set forth in Section 11(a) or Section
    11(b) unless (A) Fees have been paid in full; and (B) Company gives written notice of the defective Services,
    reasonably described, to Roofix within 7  days of the time when Company discovers or ought to have discovered
    that the Services were defective.</p>
  <p class="mb-3">(f) Subject to Section 11(e), Roofix shall, in its sole discretion, either: (i) repair or re-perform
    such Services (or the defective part); or (ii) credit or refund the price of such Services at the pro rata
    contract rate.</p>
  <p class="mb-3">THE REMEDIES SET FORTH IN SECTION 11(f) SHALL BE THE COMPANY'S SOLE AND EXCLUSIVE REMEDY AND ROOFIX'S
    ENTIRE LIABILITY FOR ANY BREACH OF THE LIMITED WARRANTIES SET FORTH IN SECTIONS 11(a) AND (b).</p>

  <h4 class="mb-3 mt-8 font-bold">12. DISCLAIMER OF WARRANTIES</h4>
  <p class="mb-3">EXCEPT FOR THE WARRANTIES SET FORTH IN SECTIONS 11(a) and 11(b) ABOVE, ROOFIX MAKES NO WARRANTY
    WHATSOEVER WITH RESPECT TO THE SERVICES, INCLUDING ANY (A) WARRANTY OF MERCHANTABILITY; OR (B) WARRANTY OF
    FITNESS FOR A PARTICULAR PURPOSE; OR (C) WARRANTY OF TITLE; OR (D) WARRANTY AGAINST INFRINGEMENT OF INTELLECTUAL
    PROPERTY RIGHTS OF A THIRD PARTY; WHETHER EXPRESS OR IMPLIED BY LAW, COURSE OF DEALING, COURSE OF PERFORMANCE,
    USAGE OF TRADE, OR OTHERWISE.</p>

  <h4 class="mb-3 mt-8 font-bold">13. LIMITATION OF LIABILITY</h4>
  <p class="mb-3">(a) IN NO EVENT SHALL ROOFIX BE LIABLE TO COMPANY OR TO ANY THIRD PARTY FOR ANY LOSS OF USE,
    REVENUE OR PROFIT OR DIMINUTION IN VALUE, OR FOR ANY CONSEQUENTIAL, INCIDENTAL, INDIRECT, EXEMPLARY, SPECIAL,
    OR PUNITIVE DAMAGES WHETHER ARISING OUT OF BREACH OF CONTRACT, TORT (INCLUDING NEGLIGENCE), OR OTHERWISE,
    REGARDLESS OF WHETHER SUCH DAMAGES WERE FORESEEABLE AND WHETHER OR NOT ROOFIX HAS BEEN ADVISED OF THE
    POSSIBILITY OF SUCH DAMAGES, AND NOTWITHSTANDING THE FAILURE OF ANY AGREED OR OTHER REMEDY OF ITS ESSENTIAL PURPOSE.</p>
  <p class="mb-3">(b) IN NO EVENT SHALL ROOFIX’ AGGREGATE LIABILITY ARISING OUT OF OR RELATED TO THIS AGREEMENT,
    WHETHER ARISING OUT OF OR RELATED TO BREACH OF CONTRACT, TORT (INCLUDING NEGLIGENCE) OR OTHERWISE, EXCEED THE
    AGGREGATE AMOUNTS PAID OR PAYABLE TO ROOFIX PURSUANT TO THE APPLICABLE PROPOSAL.</p>
  <p class="mb-3">(c) The limitation of liability set forth in Section 13(b) above shall not apply to (i) liability
    resulting from Roofix’ gross negligence or willful misconduct and (ii) death or bodily injury resulting from Roofix’
    negligent acts or omissions.</p>

  <h4 class="mb-3 mt-8 font-bold">14. TERMINATION</h4>
  <p class="mb-3">In addition to any remedies that may be provided under this Agreement, Roofix may terminate this
    Agreement with immediate effect upon written notice to Company, if Company:</p>
  <p class="mb-3">(a) fails to pay any amount when due under this Agreement and such failure continues for 5 days after
    Company's receipt of written notice of nonpayment;</p>
  <p class="mb-3">(b) has not otherwise performed or complied with any of the terms of this Agreement, in whole or in
    part, including provision to Roofix of incorrect materials selection and scope requirements; or</p>
  <p class="mb-3">(c) becomes insolvent, files a petition for bankruptcy or commences or has commenced against it
    proceedings relating to bankruptcy, receivership, reorganization, or assignment for the benefit of creditors.</p>

  <h4 class="mb-3 mt-8 font-bold">15. INSURANCE</h4>
  <p class="mb-3">During the term of this Agreement and for a period of 12 months thereafter, Company shall, at its own
    expense, maintain and carry insurance in full force and effect which includes, but is not limited to, commercial
    general liability (including product liability) in a sum sufficient to cover its obligations under this Agreement
    with financially sound and reputable insurers. Upon Roofix’ request, Company shall provide Roofix with a
    certificate of insurance from Company's insurer evidencing the insurance coverage specified in these General Terms.
    The certificate of insurance shall name Roofix as an additional insured. Company shall provide Roofix with 30 days’
    advance written notice in the event of a cancellation or material change in Company’s insurance policy. Except where
    prohibited by law, Company shall require its insurer to waive all rights of subrogation against Roofix’ insurers
    and Roofix.</p>

  <h4 class="mb-3 mt-8 font-bold">16. WAIVER</h4>
  <p class="mb-3">No waiver by Roofix of any of the provisions of this Agreement is effective unless explicitly set
    forth in writing and signed by Roofix. No failure to exercise, or delay in exercising, any rights, remedy, power,
    or privilege arising from this Agreement operates or may be construed as a waiver thereof. No single or partial
    exercise of any right, remedy, power, or privilege hereunder precludes any other or further exercise thereof or
    the exercise of any other right, remedy, power, or privilege.</p>

  <h4 class="mb-3 mt-8 font-bold">17. FORCE MAJEURE</h4>
  <p class="mb-3">Roofix shall not be liable or responsible to Company, nor be deemed to have defaulted under or
    breached this Agreement, for any failure or delay in fulfilling or performing any term of this Agreement,
    when and to the extent such failure or delay is caused by or results from acts beyond Roofix’ reasonable
    control, including, without limitation, the following force majeure events (“Force Majeure Event(s)”): (a)
    acts of God; (b) flood, fire, earthquake, other potential disasters, or catastrophes such as epidemics, or
    explosion; (c) war, invasion, hostilities (whether war is declared or not), terrorist threats or acts, riot or
    other civil unrest; (d) government order, law, or action; (e)  embargoes or blockades in effect on or after the
    date of this Agreement; (f) national or regional emergency; and (g) strikes, labor stoppages or slowdowns or other
    industrial disturbances; (h) shortage of adequate power or transportation facilities; and (i) other similar events
    beyond Roofix’ reasonable control. Roofix shall use diligent efforts to end the failure or delay and ensure the
    effects of such Force Majeure Event are minimized. Roofix shall resume the performance of its obligations as soon
    as reasonably practicable after the removal of the cause.</p>

  <h4 class="mb-3 mt-8 font-bold">18. ASSIGNMENT</h4>
  <p class="mb-3">Company shall not assign any of its rights or delegate any of its obligations under this Agreement
    without the prior written consent of Roofix. Any purported assignment or delegation in violation of this Section
    is null and void. No assignment or delegation relieves Company of any of its obligations under this Agreement.</p>

  <h4 class="mb-3 mt-8 font-bold">19. RELATIONSHIP OF THE PARTIES</h4>
  <p class="mb-3">The relationship between the parties is that of independent contractors. Nothing contained in this
    Agreement shall be construed as creating any agency, partnership, joint venture or other form of joint enterprise,
    employment, or fiduciary relationship between the parties, and neither party shall have authority to contract for
    or bind the other party in any manner whatsoever.</p>

  <h4 class="mb-3 mt-8 font-bold">20. NO THIRD-PARTY BENEFICIARIES</h4>
  <p class="mb-3">Except as expressly set forth in Section 11(b), this Agreement is for the sole benefit of the
    parties hereto and their respective successors and permitted assigns and nothing herein, express or implied, is
    intended to or shall confer upon any other person or entity any legal or equitable right, benefit or remedy of any
    nature whatsoever under or by reason of these General Terms.</p>

  <h4 class="mb-3 mt-8 font-bold">21. GOVERNING LAW; JURISDICTION</h4>
  <p class="mb-3">All matters arising out of or relating to this Agreement are governed by and construed in accordance
    with the internal laws of the State of Texas without giving effect to any choice or conflict of law provision or
    rule and shall be resolved in the Courts of Tarrant County, TX or the state courts in which the Services were
    provided.</p>

  <h4 class="mb-3 mt-8 font-bold">22. NOTICES</h4>
  <p class="mb-3">All notices, requests, consents, claims, demands, waivers, and other communications hereunder
    (each, a “Notice”) shall be in writing and addressed to the parties at the addresses set forth in the Proposal
    or to such other address that may be designated by the receiving party in writing. All Notices shall be delivered
    by personal delivery, nationally recognized overnight courier (with all fees pre-paid), email (with confirmation
    of transmission) or certified or registered mail (in each case, return receipt requested, postage prepaid).
    Except as otherwise provided in this Agreement, a Notice is effective only (a) upon receipt of the receiving party,
    and (b) if the party giving the Notice has complied with the requirements of this Section.</p>

  <h4 class="mb-3 mt-8 font-bold">23. SEVERABILITY</h4>
  <p class="mb-3">If any term or provision of this Agreement is invalid, illegal, or unenforceable in any jurisdiction,
    such invalidity, illegality, or unenforceability shall not affect any other term or provision of this Agreement or
    invalidate or render unenforceable such term or provision in any other jurisdiction.</p>

  <h4 class="mb-3 mt-8 font-bold">24. SURVIVAL</h4>
  <p class="mb-3">Provisions of these General Terms, which by their nature should apply beyond their terms, will remain
    in force after any termination or expiration of this Agreement including, but not limited to, the following provisions:
    Confidentiality, Governing Law, Insurance, Submission to Jurisdiction, and Survival.</p>

  <h4 class="mb-3 mt-8 font-bold">25. AMENDMENT AND MODIFICATION</h4>
  <p class="mb-3">A Proposal once accepted by Company may only be amended or modified in a writing which specifically
    states that it amends the applicable Proposal and is signed or expressly approved by an authorized representative of
    each party. Approval of a Change Order by Company through the Platform shall be considered express approval hereunder.</p>

  <h4 class="mb-3 mt-8 font-bold">26. CHANGES TO TERMS</h4>
  <p class="mb-3">Subject to Section 25,  Roofix may revise and update these General Terms from time to time in its
    sole discretion. All changes are effective immediately when posted.</p>

</div>
