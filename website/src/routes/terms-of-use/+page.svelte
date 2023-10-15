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
  <h1 class="text-4xl font-semibold text-center">
    RFX, Inc. Mobile App<br/>
    and Website Terms of Use
  </h1>
  <h4 class="mb-3 mt-8 font-bold">Last Updated: {updatedAt}</h4>
  <h4 class="mb-3 mt-8 font-bold">1. ACCEPTANCE OF THE TERMS OF USE</h4>
  <p class="mb-3">These terms of use (“Terms of Use”) are entered into by and between You and to RFX, Inc. (“RFX, Inc.”
    “we,” or “us”) and govern your access to and use of <span class="font-bold">{domain}</span>, and our mobile
    application (<span class="font-bold">{mobilAppName}</span>) (collectively, the website and the App are referred to
    herein as the “Site”), in addition to
    and related to information, products and services, by individuals or entities (“you” or “your”) in connection with
    roof replacement and repair estimating and construction (the “Purpose”).</p>
  <p class="mb-3">
    Please read the Terms of Use carefully before you start to use the Site. <strong>By using the Site or by clicking to
    accept or agree to the Terms of Use when this option is made available to you, you accept and agree to be bound and
    abide by these Terms of Use and our privacy policy (“Privacy Policy”), found at {termLink},
    incorporated herein by reference.</strong> If you do not want to agree to these Terms of Use or the Privacy Policy,
    you must not access or use the Site.
  </p>
  <p class="mb-3">This Site is offered and available to users who are 18 years of age or older and reside in the <span>{country}</span>.
    By using this Site, you represent and warrant that you meet all of the
    foregoing eligibility requirements. If you do not meet all of these requirements, you must not access or use the
    Site.
  </p>
  <h4 class="mb-3 mt-8 font-bold">2. CHANGES TO THE TERMS OF USE</h4>
  <p class="mb-3">We may revise and update these Terms of Use from time to time in our sole discretion. All changes are
    effective immediately when we post them and apply to all access to and use of the Site thereafter. Your continued
    use
    of the Site following the posting of revised Terms of Use means that you accept and agree to the changes. You are
    expected to check this page so you are aware of any changes, as they are binding on you.</p>
  <h4 class="mb-3 mt-8 font-bold">3. ACCESSING THE SITE AND ACCOUNT SECURITY</h4>
  <p class="mb-3">We reserve the right to withdraw or amend this Site, and any service or material we provide on the
    Site,
    in our sole discretion without notice. We will not be liable if for any reason all or any part of the Site is
    unavailable at any time or for any period.</p>
  <p class="mb-3">You are responsible for:</p>

  <ul class="list-disc mb-3 list-outside ml-6">
    <li>Making all arrangements necessary for you to have access to the Site.</li>
    <li>Ensuring that all persons who access the Site through your internet connection are aware of these Terms of Use
      and
      comply with them.
    </li>
  </ul>
  <p class="mb-3">To access the Site, you may be asked to provide certain contact information, registration details or
    other information. It is a condition of your use of the Site that all the information you provide on the Site is
    correct, current, and complete. You agree that all information you provide to register with this Site or otherwise,
    including, but not limited to, through the use of any interactive features on the Site, is governed by our Privacy
    Policy, and you consent to all actions we take with respect to your information consistent with our Privacy Policy.
    If
    you choose, or are provided with, a username, password, or any other piece of information as part of our security
    procedures, you must treat such information as confidential, and you must not disclose it to any other person or
    entity. You also acknowledge that your account is personal to you and agree not to provide any other person with
    access to this Site or portions of it using your username, password, or other security information. You agree to
    notify us immediately of any unauthorized access to or use of your username or password or any other breach of
    security. You also agree to ensure that you exit from your account at the end of each session. You should use
    particular caution when accessing your account from a public or shared computer so that others are not able to view
    or
    record your password or other personal information. We have the right to disable any username, password, or other
    identifier, whether chosen by you or provided by us, at any time in our sole discretion for any or no reason,
    including if, in our opinion, you have violated any provision of these Terms.</p>
  <h4 class="mb-3 mt-8 font-bold">4. INTELLECTUAL PROPERTY RIGHTS</h4>
  <p class="mb-3">The Site and its entire contents, features, and functionality (including but not limited to all
    information, software, text, displays, images, video, and audio, and the design, selection, and arrangement thereof)
    are owned by the RFX, Inc, its licensors, or other providers of such material and are protected by {country} and
    international copyright, trademark, patent, trade secret, and other intellectual property or proprietary rights
    laws.</p>
  <p class="mb-3">These Terms of Use permit you to use the Site for the Purpose only. Except as permitted by RFX, Inc.
    in
    connection with the Purpose, you must not reproduce, distribute, modify, create derivative works of, publicly
    display,
    publicly perform, republish, download, store, or transmit any of the material on our Site, except as follows:</p>
  <ul class="list-disc mb-3  list-outside ml-6">
    <li>Your computer may temporarily store copies of such materials incidental to your accessing and viewing those
      materials.
    </li>
    <li>You may store files that are automatically cached by your Web browser for display enhancement purposes.</li>
    <li>You may print or download one copy of a reasonable number of pages of the Site for your own personal,
      non-commercial use and not for further reproduction, publication, or distribution.
    </li>
    <li>If we provide social media features such as Facebook, Twitter, YouTube, LinkedIn and Instagram with certain
      content, you may take such actions as are enabled by such features.
    </li>
  </ul>
  <p class="mb-3">You must not:</p>
  <ul class="list-disc mb-3  list-outside ml-6">
    <li>Modify copies of any materials from the Site.</li>
    <li>Use any illustrations, photographs, video or audio sequences, or any graphics separately from the accompanying
      text.
    </li>
    <li>Delete or alter any copyright, trademark, or other proprietary rights notices from copies of materials from the
      Site.
    </li>
    <li>Except for the Purpose, you must not access or use for any commercial purposes any part of the Site or any
      services or materials available through the Site.
    </li>
  </ul>
  <p class="mb-3">If you print, copy, modify, download, or otherwise use or provide any other person with access to any
    part of the Site in breach of the Terms of Use, your right to use the Site will stop immediately and you must, at
    our
    option, return or destroy any copies of the materials you have made. No right, title, or interest in or to the Site
    or
    any content on the Site is transferred to you, and all rights not expressly granted are reserved by RFX, Inc. Any
    use
    of the Site not expressly permitted by these Terms of Use is a breach of these Terms of Use and may violate
    copyright,
    trademark, and other laws.</p>
  <h4 class="mb-3 mt-8 font-bold">5. TRADEMARKS</h4>
  <p class="mb-3">The RFX, Inc. name and all related names, logos, product and service names, designs, and slogans are
    trademarks of RFX, Inc. or its affiliates or licensors. You must not use such marks without the prior written
    permission of RFX, Inc. All other names, logos, product and service names, designs, and slogans on this Site are the
    trademarks of their respective owners.</p>
  <h4 class="mb-3 mt-8 font-bold">6. PROHIBITED USES</h4>
  <p class="mb-3">You may use the Site only for lawful purposes and in accordance with these Terms of Use. You agree not
    to use the Site in any way that violates any applicable federal, state, local, or international law or regulation
    (including, without limitation, any laws regarding the export of data or software to and from the US or other
    countries). Additionally, you agree not to:</p>
  <ul class="list-disc  list-outside ml-6">
    <li>Use the Site in any manner that could disable, overburden, damage, or impair the site or interfere with any
      other
      party's use of the Site.
    </li>
    <li>Use any robot, spider, or other automatic device, process, or means to access the Site for any purpose,
      including
      monitoring or copying any of the material on the Site.
    </li>
    <li>Use any device, software, or routine that interferes with the proper working of the Site.</li>
    <li>Introduce any viruses, Trojan horses, worms, logic bombs, or other material that is malicious or technologically
      harmful.
    </li>
    <li>Attempt to gain unauthorized access to, interfere with, damage, or disrupt any parts of the Site, the server on
      which the Site is stored, or any server, computer, or database connected to the Site.
    </li>
    <li>Attack the Site via a denial-of-service attack or a distributed denial-of-service attack.</li>
    <li>Otherwise attempt to interfere with the proper working of the Site.</li>
  </ul>
  <h4 class="mb-3 mt-8 font-bold">7. RELIANCE ON INFORMATION POSTED</h4>
  <p class="mb-3">This Site may include content provided by third parties. The information presented on or through the
    Site is made available solely for general information purposes. We do not warrant the accuracy, completeness, or
    usefulness of this information. Any reliance you place on such information is strictly at your own risk. We disclaim
    all liability and responsibility arising from any reliance placed on such materials by you or any other visitor to
    the
    Site, or by anyone who may be informed of any of its contents.</p>
  <p class="mb-3">All statements and/or opinions expressed in these materials, and all articles and responses to
    questions
    and other content, other than the content provided byRFX, Inc, are solely the opinions and the responsibility of the
    person or entity providing those materials. These materials do not necessarily reflect the opinion of RFX, Inc. We
    are
    not responsible, or liable to you or any third party, for the content or accuracy of any materials provided by any
    third parties.</p>
  <h4 class="mb-3 mt-8 font-bold">8. CHANGES TO THE SITE</h4>
  <p class="mb-3">We may update the content on this Site from time to time, but its content is not necessarily complete
    or
    up to date. Any of the material on the Site may be out of date at any given time, and we are under no obligation to
    update such material.</p>
  <h4 class="mb-3 mt-8 font-bold">9. INFORMATION ABOUT YOU AND YOUR VISITS TO THE SITE</h4>
  <p class="mb-3">All information we collect on this Site is subject to our Privacy Policy {termLink}.
    By using the Site, you consent to all actions taken by us with respect to your information in compliance with the
    Privacy Policy.</p>
  <h4 class="mb-3 mt-8 font-bold">10. LINKS FROM THE SITE</h4>
  <p class="mb-3">If the Site contains links to other sites and resources provided by third parties, these links are
    provided for your convenience only. This includes links contained in advertisements, including banner advertisements
    and sponsored links. We have no control over the contents of those sites or resources and accept no responsibility
    for
    them or for any loss or damage that may arise from your use of them. If you decide to access any of the third-party
    Sites linked to this Site, you do so entirely at your own risk and subject to the terms and conditions of use for
    such
    Sites.</p>
  <h4 class="mb-3 mt-8 font-bold">11. GEOGRAPHIC RESTRICTIONS</h4>
  <p class="mb-3">The owner of the Site is based in the State of Utah in the {country}. We provide this Site for use
    only by persons located in the <span >{country}</span>. We make no claims that the Site or
    any
    of its content is accessible or appropriate outside of the {country}. Access to the Site may not be legal by
    certain persons or in certain countries. If you access the Site from outside the {country}, you do so on your
    own
    initiative and are responsible for compliance with local laws.</p>
  <h4 class="mb-3 mt-8 font-bold">12. DISCLAIMER OF WARRANTIES</h4>
  <p class="mb-3">TO THE FULLEST EXTENT PROVIDED BY LAW, WE WILL NOT BE LIABLE FOR ANY LOSS OR DAMAGE CAUSED BY A
    DISTRIBUTED DENIAL-OF-SERVICE ATTACK, VIRUSES, OR OTHER TECHNOLOGICALLY HARMFUL MATERIAL THAT MAY INFECT YOUR
    COMPUTER
    EQUIPMENT, COMPUTER PROGRAMS, DATA, OR OTHER PROPRIETARY MATERIAL DUE TO YOUR USE OF THE SITE OR ANY SERVICES OR
    ITEMS
    OBTAINED THROUGH THE SITE OR TO YOUR DOWNLOADING OF ANY MATERIAL POSTED ON IT, OR ON ANY SITE LINKED TO IT.</p>
  <p class="mb-3">YOUR USE OF THE SITE, ITS CONTENT, AND ANY SERVICES OR ITEMS OBTAINED THROUGH THE SITE IS AT YOUR OWN
    RISK. THE SITE, ITS CONTENT, AND ANY SERVICES OR ITEMS OBTAINED THROUGH THE SITE ARE PROVIDED ON AN "AS IS" AND "AS
    AVAILABLE" BASIS, WITHOUT ANY WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED. NEITHER RFX, INC. NOR ANY PERSON
    ASSOCIATED WITH RFX, INC. MAKES ANY WARRANTY OR REPRESENTATION WITH RESPECT TO THE COMPLETENESS, SECURITY,
    RELIABILITY, QUALITY, ACCURACY, OR AVAILABILITY OF THE SITE. WITHOUT LIMITING THE FOREGOING, NEITHER RFX, INC. NOR
    ANYONE ASSOCIATED WITH RFX, INC. REPRESENTS OR WARRANTS THAT THE SITE, ITS CONTENT, OR ANY SERVICES OR ITEMS
    OBTAINED
    THROUGH THE SITE WILL BE ACCURATE, RELIABLE, ERROR-FREE, OR UNINTERRUPTED, THAT DEFECTS WILL BE CORRECTED, THAT OUR
    SITE OR THE SERVER THAT MAKES IT AVAILABLE ARE FREE OF VIRUSES OR OTHER HARMFUL COMPONENTS, OR THAT THE SITE OR ANY
    SERVICES OR ITEMS OBTAINED THROUGH THE SITE WILL OTHERWISE MEET YOUR NEEDS OR EXPECTATIONS.</p>
  <p class="mb-3">TO THE FULLEST EXTENT PROVIDED BY LAW, RFX, INC. HEREBY DISCLAIMS ALL WARRANTIES OF ANY KIND, WHETHER
    EXPRESS OR IMPLIED, STATUTORY, OR OTHERWISE, INCLUDING BUT NOT LIMITED TO ANY WARRANTIES OF MERCHANTABILITY,
    NON-INFRINGEMENT, AND FITNESS FOR PARTICULAR PURPOSE.</p>
  <p class="mb-3">THE FOREGOING DOES NOT AFFECT ANY WARRANTIES THAT CANNOT BE EXCLUDED OR LIMITED UNDER APPLICABLE
    LAW.</p>
  <h4 class="mb-3 mt-8 font-bold">13. LIMITATION ON LIABILITY</h4>
  <p class="mb-3">TO THE FULLEST EXTENT PROVIDED BY LAW, IN NO EVENT WILL RFX, INC, ITS AFFILIATES, OR THEIR LICENSORS,
    SERVICE PROVIDERS, EMPLOYEES, AGENTS, OFFICERS, OR DIRECTORS BE LIABLE FOR DAMAGES OF ANY KIND, UNDER ANY LEGAL
    THEORY, ARISING OUT OF OR IN CONNECTION WITH YOUR USE, OR INABILITY TO USE, THE SITE, ANY SITES LINKED TO IT, ANY
    CONTENT ON THE SITE OR SUCH OTHER SITES, INCLUDING ANY DIRECT, INDIRECT, SPECIAL, INCIDENTAL, CONSEQUENTIAL, OR
    PUNITIVE DAMAGES, INCLUDING BUT NOT LIMITED TO, PERSONAL INJURY, PAIN AND SUFFERING, EMOTIONAL DISTRESS, LOSS OF
    REVENUE, LOSS OF PROFITS, LOSS OF BUSINESS OR ANTICIPATED SAVINGS, LOSS OF USE, LOSS OF GOODWILL, LOSS OF DATA, AND
    WHETHER CAUSED BY TORT (INCLUDING NEGLIGENCE), BREACH OF CONTRACT, OR OTHERWISE, EVEN IF FORESEEABLE.</p>
  <h4 class="mb-3 mt-8 font-bold">14. INDEMNIFICATION</h4>
  <p class="mb-3">You agree to defend, indemnify, and hold harmlessRFX, Inc, its affiliates, licensors, and service
    providers, and its and their respective officers, directors, employees, contractors, agents, licensors, suppliers,
    successors, and assigns from and against any claims, liabilities, damages, judgments, awards, losses, costs,
    expenses,
    or fees (including reasonable attorneys' fees) arising out of or relating to your violation of these Terms of Use or
    your use of the Site, including, but not limited to, any use of the Site's content, services, and products other
    than
    as expressly authorized in these Terms of Use, or your use of any information obtained from the Site.</p>
  <h4 class="mb-3 mt-8 font-bold">15. GOVERNING LAW AND JURISDICTION</h4>
  <p class="mb-3">All matters relating to the Site and these Terms of Use, and any dispute or claim arising therefrom or
    related thereto (in each case, including non-contractual disputes or claims), shall be governed by and construed in
    accordance with the internal laws of the State of {lawState} without giving effect to any choice or conflict of law
    provision or rule. Any legal suit, action, or proceeding arising out of, or related to, these Terms of Use or the
    Site
    shall be instituted exclusively in the federal courts of the {country} or the courts of the State of Utah. You
    waive any and all objections to the exercise of jurisdiction over you by such courts and to venue in such
    courts.</p>
  <h4 class="mb-3 mt-8 font-bold">16. WAIVER AND SEVERABILITY</h4>
  <p class="mb-3">No waiver by RFX, Inc. of any term or condition set out in these Terms of Use shall be deemed a
    further
    or continuing waiver of such term or condition or a waiver of any other term or condition, and any failure of RFX,
    Inc. to assert a right or provision under these Terms of Use shall not constitute a waiver of such right or
    provision.</p>
  <p class="mb-3">If any provision of these Terms of Use is held by a court or other tribunal of competent jurisdiction
    to
    be invalid, illegal, or unenforceable for any reason, such provision shall be eliminated or limited to the minimum
    extent such that the remaining provisions of the Terms of Use will continue in full force and effect.</p>
  <h4 class="mb-3 mt-8 font-bold">17. ENTIRE AGREEMENT</h4>
  <p class="mb-3">The Terms of Use and our Privacy Policy constitute the sole and entire agreement between you and RFX,
    Inc. regarding the Site and supersede all prior and contemporaneous understandings, agreements, representations, and
    warranties, both written and oral, regarding the Site.</p>
  <h4 class="mb-3 mt-8 font-bold">18. MOBILE APPLICATION MARKETPLACES.</h4>
  <p class="mb-3">The App shall be majority deployed via the Apple AppStore and Android Marketplace. Use of the App is
    subject to terms and conditions imposed by the host environment, Apple Inc., and/or Google LLC, respectively. You
    shall reasonably cooperate with our efforts to comply with such terms. These Terms of Use are concluded between RFX,
    Inc. and, you, only (the end user), and not with Apple. As between RFX, Inc. and Apple, RFX, Inc. is solely
    responsible for the App and the content thereof. You and RFX, Inc. acknowledge that Apple has no obligation
    whatsoever
    to furnish any maintenance and support services with regard to the App. Apple, and Apple’s subsidiaries, are third
    party beneficiaries of these Terms of Use, and upon your (the end user’s) acceptance of the terms and conditions of
    these Terms of Use, Apple will have the right (and will be deemed to have accepted the right) to enforce the Terms
    or
    Use against the end user as a third party beneficiary thereof. RFX, Inc, not Apple, is solely responsible for the
    App
    and the content thereof, and RFX, Inc. shall be responsible for addressing any claims of the end user or any third
    party relating to the App or end user’s possession and/or use of the App, including, but not limited to: (i) product
    liability claims; (ii) any claim that the App fails to conform to any applicable legal or regulatory requirement;
    and
    (iii) claims arising under consumer protection, privacy, or similar legislation. In the event of any third party
    claim
    that the App or your access and/or use of the App infringes that third party’s intellectual property rights, RFX,
    Inc,
    not Apple, will be solely responsible for the investigation, defense, settlement and discharge of any such
    intellectual property infringement claim. In the event of any failure of the App to conform to any applicable
    warranty, you may notify Apple, and Apple will refund the purchase price for the App to you (if any); and, to the
    maximum extent permitted by applicable law, Apple will have no other warranty obligation whatsoever with respect to
    the App, and any other claims, losses, liabilities, damages, costs or expenses attributable to any failure to
    conform
    to any warranty will be our sole responsibility.</p>
  <h4 class="mb-3 mt-8 font-bold">YOUR COMMENTS AND CONCERNS</h4>
  <p class="mb-3">All feedback, comments, and other communications relating to the Site should be directed to: <span
    class="font-bold">{contactEmail}</span>.</p>

</div>
